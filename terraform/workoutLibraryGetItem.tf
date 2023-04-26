data "archive_file" "workoutLibraryGetItem-zip" {
  source_file = "../build/workoutLibraryGetItem"
  output_path = "../build/workoutLibraryGetItem.zip"
  type        = "zip"
}

resource "aws_lambda_function" "workoutLibraryGetItem" {
  function_name    = "workoutLibraryGetItem"
  filename         = "../build/workoutLibraryGetItem.zip"
  handler          = "workoutLibraryGetItem"
  source_code_hash = "data.archive_file.workoutLibraryGetItem-zip.output_base64sha256"
  role             = aws_iam_role.iam_for_lambda.arn
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 10
}

resource "aws_lambda_permission" "workoutLibraryGetItem" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.workoutLibraryGetItem.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.api.execution_arn}/*/*/*"
}
