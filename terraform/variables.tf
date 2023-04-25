variable "aws_region" {
  type        = string
  description = "The region in which the resources will be created"
  default     = "us-east-2"
}

variable "AWS_ACCESS_KEY_ID" {
  type        = string
  description = "The aws development account access key"
}

variable "AWS_SECRET_ACCESS_KEY" {
  type        = string
  description = "The aws development account secret key"
}

variable "app_name" {
  type        = string
  description = "The name of the application"
  default     = "plates"
}