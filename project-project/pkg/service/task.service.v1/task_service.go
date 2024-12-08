package task_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/project-common/encrypts"
	"test.com/project-common/errs"
	"test.com/project-common/tms"
	"test.com/project-grpc/task"
	"test.com/project-grpc/user/login"
	"test.com/project-project/internal/dao"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/database"
	"test.com/project-project/internal/database/tran"
	"test.com/project-project/internal/repo"
	"test.com/project-project/internal/rpc"
	"test.com/project-project/pkg/model"
	"time"
)

type TaskService struct {
	task.UnimplementedTaskServiceServer
	cache                  repo.Cache
	transaction            tran.Transaction
	projectRepo            repo.ProjectRepo
	projectTemplateRepo    repo.ProjectTemplateRepo
	taskStagesTemplateRepo repo.TaskStagesTemplateRepo
	taskStagesRepo         repo.TaskStagesRepo
	taskRepo               repo.TaskRepo
}

func New() *TaskService {
	return &TaskService{
		cache:                  dao.Rc,
		transaction:            dao.NewTransaction(),
		projectRepo:            dao.NewProjectDao(),
		projectTemplateRepo:    dao.NewProjectTemplateDao(),
		taskStagesTemplateRepo: dao.NewTaskStagesTemplateDao(),
		taskStagesRepo:         dao.NewTaskStagesDao(),
		taskRepo:               dao.NewTaskDao(),
	}
}

// grpc处理函数，所以返回的接收和返回的都是grpc的数据类型
func (t *TaskService) TaskStages(co context.Context, msg *task.TaskReqMessage) (*task.TaskStagesResponse, error) {
	projectCode := encrypts.DecryptNoErr(msg.ProjectCode)
	page := msg.Page
	pageSize := msg.PageSize

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	stages, total, err := t.taskStagesRepo.FindStagesByProjectId(ctx, projectCode, page, pageSize)
	if err != nil {
		zap.L().Error("projct TaskStages taskStagesRepo.FindStagesByProjectId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	var tsMessages []*task.TaskStagesMessage
	copier.Copy(&tsMessages, stages)
	if tsMessages == nil {
		return &task.TaskStagesResponse{List: tsMessages, Total: 0}, nil
	}
	stagesMap := data.ToTaskStagesMap(stages)
	for _, v := range tsMessages {
		taskStages := stagesMap[int(v.Id)]
		v.Code = encrypts.EncryptNoErr(int64(v.Id))
		v.CreateTime = tms.FormatByMill(taskStages.CreateTime)
		v.ProjectCode = msg.ProjectCode
	}
	return &task.TaskStagesResponse{List: tsMessages, Total: total}, nil
}

func (t *TaskService) MemberProjectList(co context.Context, msg *task.TaskReqMessage) (*task.MemberProjectResponse, error) {
	// 1.去project_member表 去查询 用户id列表
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	projectCode := encrypts.DecryptNoErr(msg.ProjectCode)
	projectMembers, total, err := t.projectRepo.FindProjectMemberByPid(ctx, projectCode)
	if err != nil {
		zap.L().Error("projct MemberProjectList projectRepo.FindProjectMemberByPid error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	// 2.拿上用户id列表 去请求用户信息
	if projectMembers == nil || len(projectMembers) <= 0 {
		return &task.MemberProjectResponse{List: nil, Total: 0}, nil
	}
	var mIds []int64
	pmMap := make(map[int64]*data.ProjectMember)
	for _, v := range projectMembers {
		mIds = append(mIds, v.MemberCode)
		pmMap[v.MemberCode] = v
	}
	// 请求用户信息
	userMsg := &login.UserMessage{
		MIds: mIds,
	}
	memberMessageList, err := rpc.LoginServiceClient.FindMemInfoByIds(ctx, userMsg)
	if err != nil {
		zap.L().Error("projct MemberProjectList LoginServiceClient.FindMemInfoByIds error", zap.Error(err))
		return nil, err
	}
	var list []*task.MemberProjectMessage
	for _, v := range memberMessageList.List {
		owner := pmMap[v.Id].IsOwner
		mpm := &task.MemberProjectMessage{
			MemberCode: v.Id,
			Name:       v.Name,
			Avatar:     v.Avatar,
			Email:      v.Email,
			Code:       v.Code,
		}
		if v.Id == owner {
			mpm.IsOwner = 1
		}
		list = append(list, mpm)
	}
	return &task.MemberProjectResponse{List: list, Total: total}, nil
}

func (t *TaskService) TaskList(ctx context.Context, msg *task.TaskReqMessage) (*task.TaskListResponse, error) {
	stageCode := encrypts.DecryptNoErr(msg.StageCode)
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	taskList, err := t.taskRepo.FindTaskByStageCode(c, int(stageCode))
	if err != nil {
		zap.L().Error("project task TaskList FindTaskByStageCode error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	var taskDisplayList []*data.TaskDisplay
	var mIds []int64
	for _, v := range taskList {
		display := v.ToTaskDisplay()
		if v.Private == 1 {
			// 隐私模式
			taskMember, err := t.taskRepo.FindTaskMemberByTaskId(ctx, v.Id, msg.MemberId)
			if err != nil {
				zap.L().Error("project task TaskList taskRepo.FindTaskMemberByTaskId error", zap.Error(err))
				return nil, errs.GrpcError(model.DBError)
			}
			if taskMember != nil {
				display.CanRead = model.CanRead
			} else {
				display.CanRead = model.NoCanRead
			}

		}
		taskDisplayList = append(taskDisplayList, display)
		mIds = append(mIds, v.AssignTo)
	}
	if mIds == nil || len(mIds) <= 0 {
		return &task.TaskListResponse{List: nil}, nil
	}
	messageList, err := rpc.LoginServiceClient.FindMemInfoByIds(ctx, &login.UserMessage{MIds: mIds})
	if err != nil {
		zap.L().Error("project task TaskList LoginServiceClient.FindMemInfoByIds error", zap.Error(err))
		return nil, err
	}
	memberMap := make(map[int64]*login.MemberMessage)
	for _, v := range messageList.List {
		memberMap[v.Id] = v
	}
	for _, v := range taskDisplayList {
		message := memberMap[encrypts.DecryptNoErr(v.AssignTo)]
		e := data.Executor{
			Name:   message.Name,
			Avatar: message.Avatar,
		}
		v.Executor = e
	}
	var taskMessageList []*task.TaskMessage
	copier.Copy(&taskMessageList, taskDisplayList)
	return &task.TaskListResponse{List: taskMessageList}, nil
}
func (t *TaskService) SaveTask(ctx context.Context, msg *task.TaskReqMessage) (*task.TaskMessage, error) {
	//先检查业务
	if msg.Name == "" {
		return nil, errs.GrpcError(model.TaskNameNotNull)
	}
	stageCode := encrypts.DecryptNoErr(msg.StageCode)
	taskStages, err := t.taskStagesRepo.FindById(ctx, int(stageCode))
	if err != nil {
		zap.L().Error("project task SaveTask taskStagesRepo.FindById error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if taskStages == nil {
		return nil, errs.GrpcError(model.TaskStagesNotNull)
	}
	projectCode := encrypts.DecryptNoErr(msg.ProjectCode)
	project, err := t.projectRepo.FindProjectById(ctx, projectCode)
	if err != nil {
		zap.L().Error("project task SaveTask projectRepo.FindProjectById error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if project == nil || project.Deleted == model.Deleted {
		return nil, errs.GrpcError(model.ProjectAlreadyDeleted)
	}
	maxIdNum, err := t.taskRepo.FindTaskMaxIdNum(ctx, projectCode)
	if err != nil {
		zap.L().Error("project task SaveTask taskRepo.FindTaskMaxIdNum error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	a := 0
	if maxIdNum == nil {
		maxIdNum = &a
	}
	maxSort, err := t.taskRepo.FindTaskSort(ctx, projectCode, stageCode)
	if err != nil {
		zap.L().Error("project task SaveTask taskRepo.FindTaskSort error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if maxSort == nil {
		maxSort = &a
	}
	assignTo := encrypts.DecryptNoErr(msg.AssignTo)
	ts := &data.Task{
		Name:        msg.Name,
		CreateTime:  time.Now().UnixMilli(),
		CreateBy:    msg.MemberId,
		AssignTo:    assignTo,
		ProjectCode: projectCode,
		StageCode:   int(stageCode),
		IdNum:       *maxIdNum + 1,
		Private:     project.OpenTaskPrivate,
		Sort:        *maxSort + 65536,
		BeginTime:   time.Now().UnixMilli(),
		EndTime:     time.Now().Add(2 * 24 * time.Hour).UnixMilli(),
	}
	err = t.transaction.Action(func(conn database.DbConn) error {
		err = t.taskRepo.SaveTask(ctx, conn, ts)
		if err != nil {
			zap.L().Error("project task SaveTask taskRepo.SaveTask error", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}
		tm := data.TaskMember{
			MemberCode: assignTo,
			TaskCode:   ts.Id,
			IsExecutor: model.Executor,
			JoinTime:   time.Now().UnixMilli(),
			IsOwner:    model.Owner,
		}
		if assignTo == msg.MemberId {
			tm.IsExecutor = model.Executor
		}
		err = t.taskRepo.SaveTaskMember(ctx, conn, tm)
		if err != nil {
			zap.L().Error("project task SaveTask taskRepo.SaveTaskMember error", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	display := ts.ToTaskDisplay()
	member, err := rpc.LoginServiceClient.FindMemInfoById(ctx, &login.UserMessage{MemId: assignTo})
	if err != nil {
		return nil, err
	}
	display.Executor = data.Executor{
		Name:   member.Name,
		Avatar: member.Avatar,
		Code:   member.Code,
	}
	tm := &task.TaskMessage{}
	copier.Copy(tm, display)
	return tm, nil
}

func (t *TaskService) TaskSort(ctx context.Context, msg *task.TaskReqMessage) (*task.TaskSortResponse, error) {
	// 移动的任务id 肯定有preTaskCode
	preTaskCode := encrypts.DecryptNoErr(msg.PreTaskCode)
	toStageCode := encrypts.DecryptNoErr(msg.ToStageCode)
	if msg.PreTaskCode == msg.NextTaskCode {
		return &task.TaskSortResponse{}, nil
	}
	err := t.sortTask(preTaskCode, msg.NextTaskCode, toStageCode)
	if err != nil {
		return nil, err
	}
	return &task.TaskSortResponse{}, nil
}

func (t *TaskService) sortTask(preTaskCode int64, nextTaskCode string, toStageCode int64) error {
	// 1.从小到大排
	// 2.重新排序后，后面的序号需要重排
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ts, err := t.taskRepo.FindTaskById(c, preTaskCode)
	if err != nil {
		zap.L().Error("project task TaskSort taskRepo.FindTaskById error", zap.Error(err))
		return errs.GrpcError(model.DBError)
	}
	err = t.transaction.Action(func(conn database.DbConn) error {
		// 如果不相等需要进行修改
		ts.StageCode = int(toStageCode)
		if nextTaskCode != "" {
			// 意味着要进行排序的替换
			DnextTaskCode := encrypts.DecryptNoErr(nextTaskCode)
			next, err := t.taskRepo.FindTaskById(c, DnextTaskCode)
			if err != nil {
				zap.L().Error("project task TaskSort taskRepo.FindTaskById error", zap.Error(err))
				return errs.GrpcError(model.DBError)
			}
			// next.Sort 要找到比它小的哪个任务
			prepre, err := t.taskRepo.FindTaskByStageCodeLtSort(c, next.StageCode, next.Sort)
			if prepre != nil {
				// 当间隙过小，直接充值sort字段的值
				if next.Sort-prepre.Sort < 50 {
					err = t.resetSort(toStageCode)
					if err != nil {
						zap.L().Error("project task TaskSort resetSort error", zap.Error(err))
						return errs.GrpcError(model.DBError)
					}
					return t.sortTask(preTaskCode, nextTaskCode, toStageCode)
				}
				ts.Sort = (prepre.Sort + next.Sort) / 2
			}
			if prepre == nil {
				ts.Sort = 0
			}
			//sort := ts.Sort
			//ts.Sort = next.Sort
			//next.Sort = sort
			//err = t.taskRepo.UpdateTaskSort(c, conn, next)
			//if err != nil {
			//	zap.L().Error("project task TaskSort taskRepo.UpdataTaskSort error", zap.Error(err))
			//	return errs.GrpcError(model.DBError)
			//}
		} else {
			maxSort, err := t.taskRepo.FindTaskSort(c, ts.ProjectCode, int64(ts.StageCode))
			if err != nil {
				zap.L().Error("project task TaskSort taskRepo.FindTaskSort error", zap.Error(err))
				return errs.GrpcError(model.DBError)
			}
			if maxSort == nil {
				a := 0
				maxSort = &a
			}
			ts.Sort = *maxSort + 65536
		}
		// 当间隙过小，直接充值sort字段的值
		//if ts.Sort < 50 {
		//	err = t.resetSort(toStageCode)
		//	if err != nil {
		//		zap.L().Error("project task TaskSort resetSort error", zap.Error(err))
		//		return errs.GrpcError(model.DBError)
		//	}
		//	return t.sortTask(preTaskCode, nextTaskCode, toStageCode)
		//}
		err = t.taskRepo.UpdateTaskSort(c, conn, ts)
		if err != nil {
			zap.L().Error("project task TaskSort taskRepo.UpdataTaskSort error", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}

		return nil
	})
	return err
}

func (t *TaskService) resetSort(stageCode int64) error {
	list, err := t.taskRepo.FindTaskByStageCode(context.Background(), int(stageCode))
	if err != nil {
		return err
	}
	return t.transaction.Action(func(conn database.DbConn) error {
		iSort := 65536
		for index, v := range list {
			v.Sort = (index + 1) * iSort
			err := t.taskRepo.UpdateTaskSort(context.Background(), conn, v)
			if err != nil {
				return err
			}
		}
		return nil
	})

}

func (t *TaskService) MyTaskList(ctx context.Context, msg *task.TaskReqMessage) (*task.MyTaskListResponse, error) {
	var tsList []*data.Task
	var err error
	var total int64
	if msg.TaskType == 1 {
		//我执行的
		tsList, total, err = t.taskRepo.FindTaskByAssignTo(ctx, msg.MemberId, int(msg.Type), msg.Page, msg.PageSize)
		if err != nil {
			zap.L().Error("project task MyTaskList taskRepo.FindTaskByAssignTo error", zap.Error(err))
			return nil, errs.GrpcError(model.DBError)
		}
	}
	if msg.TaskType == 2 {
		//我参与的
		tsList, total, err = t.taskRepo.FindTaskByMemberCode(ctx, msg.MemberId, int(msg.Type), msg.Page, msg.PageSize)
		if err != nil {
			zap.L().Error("project task MyTaskList taskRepo.FindTaskByMemberCode error", zap.Error(err))
			return nil, errs.GrpcError(model.DBError)
		}
	}
	if msg.TaskType == 3 {
		//我创建的
		tsList, total, err = t.taskRepo.FindTaskByCreateBy(ctx, msg.MemberId, int(msg.Type), msg.Page, msg.PageSize)
		if err != nil {
			zap.L().Error("project task MyTaskList taskRepo.FindTaskByCreateBy error", zap.Error(err))
			return nil, errs.GrpcError(model.DBError)
		}
	}
	if tsList == nil || len(tsList) <= 0 {
		return &task.MyTaskListResponse{List: nil, Total: 0}, nil
	}
	var pids []int64
	var mids []int64
	for _, v := range tsList {
		pids = append(pids, v.ProjectCode)
		mids = append(mids, v.AssignTo)
	}
	pList, err := t.projectRepo.FindProjectByIds(ctx, pids)
	projectMap := data.ToProjectMap(pList)

	mList, err := rpc.LoginServiceClient.FindMemInfoByIds(ctx, &login.UserMessage{
		MIds: mids,
	})
	mMap := make(map[int64]*login.MemberMessage)
	for _, v := range mList.List {
		mMap[v.Id] = v
	}
	var mtdList []*data.MyTaskDisplay
	for _, v := range tsList {
		memberMessage := mMap[v.AssignTo]
		name := memberMessage.Name
		avatar := memberMessage.Avatar
		mtd := v.ToMyTaskDisplay(projectMap[v.ProjectCode], name, avatar)
		mtdList = append(mtdList, mtd)
	}
	var myMsgs []*task.MyTaskMessage
	copier.Copy(&myMsgs, mtdList)
	return &task.MyTaskListResponse{List: myMsgs, Total: total}, nil
}
