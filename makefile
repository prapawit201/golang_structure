# กำหนดที่อยู่ของ mockgen
MOCKGEN = mockgen

# กำหนด source และ destination
SOURCE = user/service/init.go
DEST = user/service/mocks/user.go

# คำสั่ง generate mocks
mockgen:
	$(MOCKGEN) -source=$(SOURCE) -destination=$(DEST)

.PHONY: mockgen
