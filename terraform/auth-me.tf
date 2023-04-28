data "archive_file" "authMe-zip" {
  source_file = "../build/authMe"
  output_path = "../build/authMe.zip"
  type        = "zip"
}

resource "aws_lambda_function" "authMe" {
  function_name    = "authMe"
  filename         = "../build/authMe.zip"
  handler          = "authMe"
  source_code_hash = "data.archive_file.authMe-zip.output_base64sha256"
  role             = aws_iam_role.iam_for_lambda.arn
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 10
}

resource "aws_lambda_permission" "authMe" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.authMe.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.api.execution_arn}/*/*/*"
}
