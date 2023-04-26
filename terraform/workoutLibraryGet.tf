data "archive_file" "workoutLibraryGet-zip" {
  source_file = "../build/workoutLibraryGet"
  output_path = "../build/workoutLibraryGet.zip"
  type        = "zip"
}

resource "aws_lambda_function" "workoutLibraryGet" {
  function_name    = "workoutLibraryGet"
  filename         = "../build/workoutLibraryGet.zip"
  handler          = "workoutLibraryGet"
  source_code_hash = "data.archive_file.workoutLibraryGet-zip.output_base64sha256"
  role             = aws_iam_role.iam_for_lambda.arn
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 10
}

resource "aws_lambda_permission" "workoutLibraryGet" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.workoutLibraryGet.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.api.execution_arn}/*/*/*"
}
