# ShopLine Order Application

The project simulates features, purchases, and deliveries

## Overview
- Using JWT for authentication
- Apply Clean Architecture for project (Learn more at [200lab.io/blog/ung-dung-clean-architecture-service-golang-rest-api](https://200lab.io/blog/ung-dung-clean-architecture-service-golang-rest-api/))



## Technology stack
- Gin
- GORM

### Prerequisites
Before you begin, you'll need to have Go installed on your computer. You can download Go from the [official Go website](https://golang.org/dl/).

### Installation
1. Clone the project from GitHub:

```bash
git clone https://github.com/tungdevpro/coffee_api.git
```

2. Add .env file
```env
PORT=3002
HASH_SALT=your_hash_salt
DB_CONNECTION_URL="your_db_url"
SECRET_KEY=your_secret_key
S3_BUCKET_NAME="your_bucket_name"
S3_REGION="ap-southeast-1"
S3_API_KEY="your_api_key"
s3_SECRET_KEY="your_s3_secret_key"
S3_DOMAIN="your_domain_s3"
```

3. For Docker
```bash
docker compose up -d
```

4. Run application by command:
```bash
go run cmd/main.go
```

5. Postman (Optional)
Download link API PostMan
