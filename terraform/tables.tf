resource "aws_dynamodb_table" "tf_notes_table" {
  name         = "tf-notes-table"
  billing_mode = "PAY_PER_REQUEST"
  attribute {
    name = "noteId"
    type = "S"
  }
  hash_key = "noteId"
}