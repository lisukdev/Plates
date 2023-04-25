resource "aws_dynamodb_table" "workout_templates" {
  name         = "WorkoutTemplates"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"

  attribute {
    name = "id"
    type = "S"
  }
}

resource "aws_dynamodb_table" "user_workout_templates" {
  name         = "UserWorkoutTemplates"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "userId"
  range_key    = "templateId"

  attribute {
    name = "userId"
    type = "S"
  }
  attribute {
    name = "templateId"
    type = "S"
  }
}
