resource "aws_dynamodb_table" "workout_templates" {
  name         = "WorkoutTemplates"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "Id"

  attribute {
    name = "Id"
    type = "S"
  }
}

resource "aws_dynamodb_table" "workout_template_libraries" {
  name         = "WorkoutTemplateLibraries"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LibraryId"
  range_key    = "TemplateId"

  attribute {
    name = "LibraryId"
    type = "S"
  }
  attribute {
    name = "TemplateId"
    type = "S"
  }
}
