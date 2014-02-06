package repository

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/pa001024/MoeFeed/app/models"
	"github.com/robfig/revel"
)

func PlatformRepo() *Platform { return &Platform{QbsRepo()} }

// 多对一关联Account
type Platform struct{ *QbsRepository }

// 聚集索引
func (this *Platform) GetUser(id interface{}) (m *models.PlatformUser) {
	this.GetRef(&m, "platform_user.id", id)
	return
}

// 索引
func (this *Platform) FindUser(account_id interface{}) (m []*models.PlatformUser) {
	this.Find(&m, "account_id", account_id)
	return
}

// 索引
func (this *Platform) FindByName(name string) (m []*models.PlatformUser) {
	this.Find(&m, "name", name)
	return
}

// 二级索引 左连接
func (this *Platform) FindByUsername(username string) (m []*models.PlatformUser) {
	this.Find(&m, "account.username", username)
	return
}

// url获取
func (this *Platform) GetProject(userName, projectName string) (m *models.Project) {
	this.GetNRef(&m, "platform_user.name = ?", userName, "project.name", projectName)
	return
}

// 列出用户所有项目
func (this *Platform) FindProjectByOwner(owner_id int64) (m []*models.Project) {
	this.OrderByDesc("updated")
	this.Find(&m, "owner_id", owner_id)
	return
}

// 列出用户所有项目
func (this *Platform) FindByOwnerPublic(owner_id int64) (m []*models.Project) {
	this.OrderByDesc("updated")
	this.FindN(&m, 0, "owner_id = ?", owner_id, "type = ?", models.ProjectPublic)
	return
}

// 或获取特定用户对项目的访问权限
func (this *Platform) GetAccess(user_id int64, project_id int64) (m *models.ProjectAccess) {
	this.OmitJoin()
	this.FindN(&m, 0, "user_id = ?", user_id, "project_id = ?", project_id)
	return
}

// 保存并储存资源
func (this *Platform) PutAndStoneResource(resource *models.Resource, r multipart.File) {
	defer r.Close()
	// 计算Hash
	m := md5.New()
	r.Seek(0, 0)
	io.Copy(m, r)
	resource.Hash = fmt.Sprintf("%x", m.Sum(nil))
	fn := resource.FileName()
	os.MkdirAll(fn[:len(fn)-32], 0644)
	f, err := os.Create(fn)
	if err != nil {
		revel.ERROR.Printf("ProjectRepo.PutAndStone() throws %v", err)
		return
	}
	defer f.Close()
	r.Seek(0, 0)
	size, _ := io.Copy(f, r)
	resource.Size = size
	this.Put(resource)
}

// 获取文件
func (this *Platform) GetResourceFile(resource *models.Resource) io.ReadCloser {
	// TODO: 缓存?
	f, err := os.Open(resource.FileName())
	if err != nil {
		revel.ERROR.Printf("ProjectRepo.GetFile() throws %v", err)
	}
	return f
}

// 联合聚集索引
func (this *Platform) GetResource(name string, project_id int64) (m *models.Resource) {
	this.GetNRef(&m, "name = ?", name, "project_id", project_id)
	return
}

// 列出项目所有资源
func (this *Platform) FindResource(project_id int64) (m []*models.Resource) {
	this.Find(&m, "project_id", project_id)
	return
}

// 联合聚集索引
func (this *Platform) GetSource(name string, project_id int64) (m *models.Source) {
	this.GetNRef(&m, "name = ?", name, "project_id = ?", project_id)
	return
}

// 列出项目所有Source
func (this *Platform) FindSource(project_id int64) (m []*models.Source) {
	this.Find(&m, "project_id", project_id)
	return
}

// 联合聚集索引
func (this *Platform) GetTarget(name string, project_id int64) (m *models.Target) {
	this.GetNRef(&m, "name = ?", name, "project_id = ?", project_id)
	return
}

// 列出项目所有Target
func (this *Platform) FindTarget(project_id int64) (m []*models.Target) {
	this.Find(&m, "project_id", project_id)
	return
}

// 联合聚集索引
func (this *Platform) GetFilter(name string, project_id int64) (m *models.Filter) {
	this.GetNRef(&m, "name = ?", name, "project_id = ?", project_id)
	return
}

// 列出项目所有Filter
func (this *Platform) FindFilter(project_id int64) (m []*models.Filter) {
	this.Find(&m, "project_id", project_id)
	return
}

// 列出项目所有Channel
func (this *Platform) GetChannel(name string, project_id int64) (m *models.Channel) {
	this.GetNRef(&m, "name = ?", name, "project_id = ?", project_id)
	return
}

// 列出项目所有Channel
func (this *Platform) FindChannel(project_id int64) (m []*models.Channel) {
	this.Find(&m, "project_id", project_id)
	return
}

// 联合聚集索引
func (this *Platform) GetCallback(url string, project_id int64) (m *models.Callback) {
	this.GetNRef(&m, "url = ?", url, "project_id = ?", project_id)
	return
}

// 列出项目所有Callback
func (this *Platform) FindCallback(project_id int64) (m []*models.Callback) {
	this.Find(&m, "project_id", project_id)
	return
}

// 列出项目所有UserActivity
func (this *Platform) FindProjectUserActivity(project_id int64) (m []*models.UserActivity) {
	this.Find(&m, "project_id", project_id)
	return
}

// 列出用户项目所有UserActivity
func (this *Platform) FindUserActivityWithUser(user_id, project_id int64) (m []*models.UserActivity) {
	this.FindN(&m, 0, "user_id = ?", user_id, "project_id = ?", project_id)
	return
}

// 列出用户所有UserActivity
func (this *Platform) GetUserActivity(id int64) (m *models.UserActivity) {
	m = &models.UserActivity{}
	this.Find(&m, "id", id)
	return
}

// 列出用户所有UserActivity
func (this *Platform) FindUserActivity(user_id int64) (m []*models.UserActivity) {
	this.Find(&m, "user_id", user_id)
	return
}
