data "archive_file" "function-zip" {
  source_file = "../build/function"
  output_path = "../build/function.zip"
  type        = "zip"
}

resource "aws_lambda_function" "function" {
  function_name    = "function"
  filename         = "../build/function.zip"
  handler          = "function"
  source_code_hash = "data.archive_file.function-zip.output_base64sha256"
  role             = aws_iam_role.iam_for_lambda.arn
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 10
}

resource "aws_lambda_permission" "apigw_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.function.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.api.execution_arn}/*/*/*"
}
