package task

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/lib4go/security/md5"
)

//Task cron任务的task明细
type Task struct {
	Cron    string `json:"cron" valid:"ascii,required" toml:"cron,omitempty"`
	Service string `json:"service" valid:"ascii,required" toml:"service,omitempty"`
	Disable bool   `json:"disable,omitemptye" toml:"disable,omitempty"`
}

//NewTask 创建任务信息
func NewTask(cron string, service string) *Task {
	return &Task{
		Cron:    cron,
		Service: service,
	}
}

//GetUNQ 获取任务的唯一标识
func (t *Task) GetUNQ() string {
	return md5.Encrypt(fmt.Sprintf("%s(%s)", t.Service, t.Cron))
}

//Validate 验证任务参数
func (t *Task) Validate() error {
	if b, err := govalidator.ValidateStruct(t); !b && err != nil {
		return fmt.Errorf("task配置有误:%v", err)
	}
	return nil
}