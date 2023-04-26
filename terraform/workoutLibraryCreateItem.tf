data "archive_file" "workoutLibraryCreateItem-zip" {
  source_file = "../build/workoutLibraryCreateItem"
  output_path = "../build/workoutLibraryCreateItem.zip"
  type        = "zip"
}

resource "aws_lambda_function" "workoutLibraryCreateItem" {
  function_name    = "workoutLibraryCreateItem"
  filename         = "../build/workoutLibraryCreateItem.zip"
  handler          = "workoutLibraryCreateItem"
  source_code_hash = "data.archive_file.workoutLibraryCreateItem-zip.output_base64sha256"
  role             = aws_iam_role.iam_for_lambda.arn
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 10
}

resource "aws_lambda_permission" "workoutLibraryCreateItem" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_workoutLibraryCreateItem.workoutLibraryCreateItem.workoutLibraryCreateItem_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.api.execution_arn}/*/*/*"
}
