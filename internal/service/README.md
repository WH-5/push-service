# Service
客户端发送的消息格式：
{
"type": 1,
"target_unique": "man",        // 接收者 ID
"payload": {
"content": "你好，世界",
"message_type": "text"  
}
}

服务端转发的格式：
{
"type": 1,
"payload": "eyJjb250ZW50Ijoi5L2g5aW9IiwibWVzc2FnZV90eXBlIjoidGV4dCJ9",
"unique_id": "what",         //发送者 ID
"timestamp": "2025-05-08 13:00:00"
}












