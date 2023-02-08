#!/bin/bash

pkg_name="$(echo -e "${1}" | tr -d '[:space:]')"
echo "$pkg_name"
if [ ! -d "./pkg/$pkg_name" ]; then
	echo "creating $pkg_name"
  	mkdir -p "./pkg/$pkg_name";
else
	echo "$pkg_name pkg already exists"
	exit
fi

echo "package $pkg_name

import (
	\"gorm.io/gorm\"
)

type repo struct {
	DB *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}
" > "./pkg/$pkg_name/postgres.go"

echo "package $pkg_name

type Repository interface {
}
" > "./pkg/$pkg_name/repository.go"


echo "package $pkg_name


type Service interface {

}

type ${1}Svc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &${1}Svc{repo: r}
}
"  > "./pkg/$pkg_name/service.go"