resource "aws_dynamodb_table" "workout_templates" {
  name         = "WorkoutTemplates"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "Id"

  attribute {
    name = "Id"
    type = "S"
  }
}

resource "aws_dynamodb_table" "user_workout_templates" {
  name         = "UserWorkoutTemplates"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "UserId"
  range_key    = "TemplateId"

  attribute {
    name = "UserId"
    type = "S"
  }
  attribute {
    name = "TemplateId"
    type = "S"
  }
}
