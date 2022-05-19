use chrono::DateTime;

pub struct Item {
    id: i32,
    name: String,
    is_lent_item: bool,
    img_uri: String,
    lend_start: DateTime,
    lend_end: DateTime
}