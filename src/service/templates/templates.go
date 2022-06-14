package templates

import (
	"encoding/json"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/repository"
	"github.com/sirupsen/logrus"
)

type TemplatesService interface {
	Get(name string) (*entity.Templates, error)
	GetById(id string) (*entity.Templates, error)
	Page(num, size int) (*entity.Page, error)
	List() ([]entity.Templates, error)
	Save(template entity.Templates) error
	Update(template entity.Templates) error
	Delete(name string) error
	Batch(templates entity.OperateTemplates) error
}

type templatesService struct {
	templateRepo repository.TemplatesRepository
}

func NewTemplatesService() TemplatesService {
	return &templatesService{
		templateRepo: repository.NewTemplatesRepository(),
	}
}

func (t templatesService) Get(name string) (*entity.Templates, error) {
	var template entity.Templates
	modelTemplate, err := t.templateRepo.Get(name)
	if err != nil {
		return nil, err
	}
	template.Templates = modelTemplate
	return &template, nil
}

func (t templatesService) GetById(id string) (*entity.Templates, error) {
	var template entity.Templates
	modelTemplate, err := t.templateRepo.Get(id)
	if err != nil {
		return nil, err
	}
	template.Templates = modelTemplate
	return &template, nil
}

func (t templatesService) Page(num, size int) (*entity.Page, error) {
	total, templates, err := t.templateRepo.Page(num, size)
	if err != nil {
		return nil, err
	}
	page := entity.Page{}
	page.Total = total
	page.Items = templates
	return &page, nil
}

func (t templatesService) List() ([]entity.Templates, error) {
	templates, err := t.templateRepo.List()
	if err != nil {
		return nil, err
	}
	var templatesEntity []entity.Templates
	for _, template := range templates {
		var templateEntity entity.Templates
		templateEntity.Templates = template
		templatesEntity = append(templatesEntity, templateEntity)
	}
	return templatesEntity, nil
}

func (t templatesService) Save(template entity.Templates) error {
	loginfo, _ := json.Marshal(template)
	logger.Log.WithFields(logrus.Fields{"template_info": string(loginfo)}).Debugf("start to add the template %s", template.Name)
	modelTemplate, err := t.templateRepo.Get(template.Name)
	if err != nil {
		return err
	}
	if &modelTemplate != nil {
		return constant.ErrResourceExist
	}
	err = t.templateRepo.Save(&template.Templates)
	if err != nil {
		return err
	}
	return nil
}

func (t templatesService) Update(template entity.Templates) error {
	jsonTemplate, err := json.Marshal(template)
	if err != nil {
		return err
	}

	mapTemplate := make(map[string]interface{})

	err = json.Unmarshal(jsonTemplate, &mapTemplate)
	if err != nil {
		return err
	}
	return t.templateRepo.Update(template.ID, mapTemplate)
}

func (t templatesService) Delete(name string) error {
	return t.templateRepo.Delete(name)
}

func (t templatesService) Batch(templates entity.OperateTemplates) error {
	return t.templateRepo.Batch(templates.Operation, templates.Items)
}
