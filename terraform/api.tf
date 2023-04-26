variable "api_authorization" {
  type        = string
  description = "The authorization scheme used on all api methods"
  default     = "NONE"
}

data "template_file" "apidef" {
  template = file("../openapi.yaml")
  vars = {
    aws_region = "us-east-2"

    function_arn = aws_lambda_function.function.arn
  }
}

resource "aws_api_gateway_rest_api" "api" {
  name = "PlatesApi"
  body = data.template_file.apidef.rendered
}
