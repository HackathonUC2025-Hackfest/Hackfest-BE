package storage

import supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"

func New(projectUrl string, token string, bucketName string) *supabasestorageuploader.Client {
	return supabasestorageuploader.New(projectUrl, token, bucketName)
}
