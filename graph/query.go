package graph

import (
	"context"
	"fmt"
	"net/url"
	"nft-marketplace/entity"
	"nft-marketplace/helper"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
)

var (
	uni *ut.UniversalTranslator
)

type itemQuery struct {
	entity.ItemQuery
}

type itemInput struct {
	entity.ItemInput
}

type itemUpdate struct {
	entity.ItemUpdate
}

func NewItemQuery(filter *Filter) itemQuery {
	i := itemQuery{}

	if filter != nil {
		i.ItemQuery = entity.ItemQuery{
			Rating:          filter.Rating,
			ReputationBadge: filter.ReputationBadge,
			Category:        filter.Category,
		}

		if filter.Availability != nil {
			i.ItemQuery.Availability = &entity.RangeInput{
				Gte: filter.Availability.Gte,
				Lte: filter.Availability.Lte,
			}
		}
	}

	return i
}

func NewItemInput(input CreateItem) itemInput {
	i := itemInput{
		ItemInput: entity.ItemInput{
			Name:         input.Name,
			Rating:       input.Rating,
			Category:     input.Category,
			Image:        input.Image,
			Reputation:   input.Reputation,
			Price:        input.Price,
			Availibility: input.Availibility,
		},
	}

	return i
}

func NewItemUpdate(input UpdateItem) itemUpdate {
	i := itemUpdate{
		ItemUpdate: entity.ItemUpdate{
			Name:         input.Name,
			Rating:       input.Rating,
			Category:     input.Category,
			Image:        input.Image,
			Reputation:   input.Reputation,
			Price:        input.Price,
			Availibility: input.Availibility,
		},
	}

	return i
}

func (item *itemInput) Validate(ctx context.Context) validator.ValidationErrors {
	en := en.New()
	uni = ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()

	validate.RegisterValidation("longer_10", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		if len(val) < 10 {
			return false
		}
		return true
	})

	validate.RegisterValidation("word_alert", func(fl validator.FieldLevel) bool {
		keywords := []string{"sex", "gay", "lesbian"}

		p := "(?i:(" + strings.Join(keywords, ")|(") + "))"
		val := fl.Field().String()

		re := regexp.MustCompile(p)

		if re.MatchString(val) {
			return false
		}

		return true
	})

	validate.RegisterValidation("0_5", func(fl validator.FieldLevel) bool {
		val := fl.Field().Int()
		if val >= 0 && val <= 5 {
			return true
		}

		return false
	})

	validate.RegisterValidation("0_1000", func(fl validator.FieldLevel) bool {
		val := fl.Field().Int()
		if val >= 0 && val <= 1000 {
			return true
		}

		return false
	})

	validate.RegisterValidation("category", func(fl validator.FieldLevel) bool {
		keywords := []string{"photo", "sketch", "cartoon", "animation"}
		val := fl.Field().String()

		return helper.Contains(keywords, val)
	})

	validate.RegisterValidation("url", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		if _, err := url.ParseRequestURI(val); err != nil {
			return false
		}

		return true
	})

	validate.RegisterTranslation("longer_10", trans, func(ut ut.Translator) error {
		return ut.Add("longer_10", "must be longer than 10", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("longer_10", fe.Field())

		return strings.ToLower(t)
	})

	validate.RegisterTranslation("word_alert", trans, func(ut ut.Translator) error {
		return ut.Add("word_alert", "name contain forbidden word", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("word_alert", fe.Field())

		return strings.ToLower(t)
	})

	validate.RegisterTranslation("0_5", trans, func(ut ut.Translator) error {
		return ut.Add("0_5", "must be higher equal than 0 and lower equal than 5", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("0_5", fe.Field())

		return strings.ToLower(t)
	})

	validate.RegisterTranslation("0_1000", trans, func(ut ut.Translator) error {
		return ut.Add("0_1000", "must be higher equal than 0 and lower equal than 1000", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("0_1000", fe.Field())

		return strings.ToLower(t)
	})

	validate.RegisterTranslation("category", trans, func(ut ut.Translator) error {
		return ut.Add("category", "must be one of [photo, sketch, cartoon, animation]", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("category", fe.Field())

		return strings.ToLower(t)
	})

	validate.RegisterTranslation("url", trans, func(ut ut.Translator) error {
		return ut.Add("url", "must be valid url", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("url", fe.Field())

		return strings.ToLower(t)
	})

	if err := validate.Struct(item.ItemInput); err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)

		resultErr := []map[string]interface{}{}

		for _, e := range errs {
			fieldName := e.StructField()
			t := reflect.TypeOf(item.ItemInput)
			field, _ := t.FieldByName(fieldName)
			j := field.Tag.Get("json")

			temp := map[string]interface{}{
				"field":   j,
				"message": e.Translate(trans),
			}

			resultErr = append(resultErr, temp)
		}

		AddInputError(ctx, resultErr)

		return errs
	}

	return nil
}

func (item *itemUpdate) Validate(ctx context.Context) validator.ValidationErrors {
	en := en.New()
	uni = ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()

	validate.RegisterValidation("longer_10", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()

		if len(val) < 10 {
			return false
		}
		return true
	})

	validate.RegisterValidation("word_alert", func(fl validator.FieldLevel) bool {
		keywords := []string{"sex", "gay", "lesbian"}

		p := "(?i:(" + strings.Join(keywords, ")|(") + "))"
		val := fl.Field().String()

		re := regexp.MustCompile(p)

		if re.MatchString(val) {
			return false
		}

		return true
	})

	validate.RegisterValidation("0_5", func(fl validator.FieldLevel) bool {
		val := fl.Field().Int()

		if val >= 0 && val <= 5 {
			return true
		}

		return false
	})

	validate.RegisterValidation("0_1000", func(fl validator.FieldLevel) bool {
		val := fl.Field().Int()
		if val >= 0 && val <= 1000 {
			return true
		}

		return false
	})

	validate.RegisterValidation("category", func(fl validator.FieldLevel) bool {
		keywords := []string{"photo", "sketch", "cartoon", "animation"}
		val := fl.Field().String()
		fmt.Println("val", val)
		return helper.Contains(keywords, val)
	})

	validate.RegisterValidation("url", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		fmt.Println("val", val)
		if _, err := url.ParseRequestURI(val); err != nil {
			return false
		}

		return true
	})

	validate.RegisterTranslation("longer_10", trans, func(ut ut.Translator) error {
		return ut.Add("longer_10", "must be longer than 10", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("longer_10", fe.Field())

		return strings.ToLower(t)
	})

	validate.RegisterTranslation("word_alert", trans, func(ut ut.Translator) error {
		return ut.Add("word_alert", "name contain forbidden word", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("word_alert", fe.Field())

		return strings.ToLower(t)
	})

	validate.RegisterTranslation("0_5", trans, func(ut ut.Translator) error {
		return ut.Add("0_5", "must be higher equal than 0 and lower equal than 5", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("0_5", fe.Field())

		return strings.ToLower(t)
	})

	validate.RegisterTranslation("0_1000", trans, func(ut ut.Translator) error {
		return ut.Add("0_1000", "must be higher equal than 0 and lower equal than 1000", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("0_1000", fe.Field())

		return strings.ToLower(t)
	})

	validate.RegisterTranslation("category", trans, func(ut ut.Translator) error {
		return ut.Add("category", "must be one of [photo, sketch, cartoon, animation]", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("category", fe.Field())

		return strings.ToLower(t)
	})

	validate.RegisterTranslation("url", trans, func(ut ut.Translator) error {
		return ut.Add("url", "must be valid url", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("url", fe.Field())

		return strings.ToLower(t)
	})

	if err := validate.Struct(item.ItemUpdate); err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)

		resultErr := []map[string]interface{}{}

		for _, e := range errs {
			fieldName := e.StructField()
			t := reflect.TypeOf(item.ItemUpdate)
			field, _ := t.FieldByName(fieldName)
			j := field.Tag.Get("json")

			temp := map[string]interface{}{
				"field":   j,
				"message": e.Translate(trans),
			}

			resultErr = append(resultErr, temp)
		}

		AddInputError(ctx, resultErr)

		return errs
	}

	return nil
}
