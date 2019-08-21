package service

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
)

type TagService struct {
	repository     *repository.Factory
	validate       *validator.Validate
	tagRepository *repository.TagRepository
}

func AddTagsToPhoto()  {
	var tags []string
	tags = append(tags, "Foo", "Bar")
	fmt.Println(strings.Join(tags, ", "))
}
