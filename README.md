
### คำอธิบาย:
- **Health Check API**: เพิ่ม endpoint `/healthcheck` สำหรับเช็คสถานะของเซิร์ฟเวอร์
    - **Method**: `GET`
    - **Response**: คืนค่า `{"status": "ok"}` หากเซิร์ฟเวอร์ทำงานปกติ
    - **ตัวอย่างการใช้งาน**: ใช้คำสั่ง `curl` เพื่อทดสอบ API นี้
- **โครงสร้างโปรเจกต์**: โครงสร้างไฟล์และโฟลเดอร์ต่างๆ ที่ใช้ในโปรเจกต์

### การใช้งาน:
- หลังจากที่คุณรันแอปพลิเคชันด้วยคำสั่ง `go run main.go` แล้ว สามารถเช็คสถานะของเซิร์ฟเวอร์ได้ที่ `http://localhost:8080/healthcheck`
  
หากต้องการการปรับเปลี่ยนใดๆ เพิ่มเติมสามารถบอกได้นะครับ!
