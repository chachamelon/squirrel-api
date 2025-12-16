# API Test Engine

간단하고 강력한 API 테스트 자동화 도구입니다. YAML 파일로 테스트 케이스를 정의하고 HTTP 요청을 실행하여 응답을 검증합니다.

## 특징

- YAML 기반의 간단한 테스트 케이스 정의
- HTTP 요청 및 응답 검증
- JSON 응답의 특정 필드 검증
- 빠르고 가벼운 Go 기반 구현

## 요구사항

- Go 1.19 이상

## 설치

### 1. 저장소 클론
```bash
git clone <repository-url>
cd api-test-engine
```

### 2. 의존성 설치
```bash
go mod tidy
```

이 명령어는 `go.mod`와 `go.sum` 파일에 정의된 모든 의존성을 자동으로 다운로드합니다.

### 3. 빌드
```bash
go build -o runner cmd/runner/main.go
```

또는 전체 모듈을 빌드하려면:
```bash
go build ./...
```

## 사용법

### 기본 사용법

```bash
# 단일 테스트 파일 실행
./runner testcases/sample.yaml

# 디렉토리 내 모든 테스트 파일 실행
./runner testcases/
```

### 테스트 케이스 형식

YAML 파일로 테스트 케이스를 정의합니다:

```yaml
name: login_api                    # 테스트 케이스 이름
request:
  method: GET                      # HTTP 메소드 (GET, POST, PUT, DELETE 등)
  url: "https://api.example.com/login"  # 요청 URL
  headers:                         # 요청 헤더 (선택사항)
    Content-Type: application/json
    Authorization: Bearer token
  body:                            # 요청 본문 (선택사항)
    username: testuser
    password: testpass
assert:
  status: 200                      # 예상 HTTP 상태 코드
  json:                            # JSON 응답 검증 (선택사항)
    token: not_null                # 필드가 존재하고 null이 아님을 검증
    user.id: not_null              # 중첩된 JSON 필드 검증도 지원
```

### 검증 옵션

- `not_null`: 필드가 존재하고 null이 아님을 검증
- `null`: 필드가 null임을 검증
- 또는 특정 값과 일치하는지 검증

## 예제

프로젝트의 `testcases/` 디렉토리에 샘플 테스트 케이스가 포함되어 있습니다.

```bash
# 샘플 테스트 실행
./runner testcases/sample.yaml
```

## 프로젝트 구조

```
.
├── cmd/runner/          # 메인 실행 파일
├── internal/
│   ├── assert/         # 응답 검증 로직
│   ├── executor/       # 테스트 실행 엔진
│   ├── model/          # 데이터 모델
│   └── parser/         # YAML 파서
├── testcases/          # 샘플 테스트 케이스
├── go.mod             # Go 모듈 정의 파일
├── go.sum             # 의존성 체크섬 파일
├── .gitignore         # Git 무시 파일
└── README.md
```

### 주요 파일 설명

- **`go.mod`**: Go 모듈의 이름, Go 버전, 의존성을 정의
- **`go.sum`**: 다운로드된 의존성 패키지의 체크섬을 저장하여 보안을 보장
- **`cmd/runner/main.go`**: 애플리케이션의 진입점

## 라이선스

이 프로젝트는 MIT 라이선스 하에 배포됩니다.
