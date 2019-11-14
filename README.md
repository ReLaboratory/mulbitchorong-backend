# mulbitchorong-backend

안드로이드 앱 물빛초롱의 서버단 Repository입니다.

## API
=> 직렬화 포맷은 JSON입니다.

- Account API
  - `POST api/account/login` : 로그인 정보를 받으면 사용자이름과 로그인 성공 여부를 응답합니다.
    - request header: X
    - params: X
    - request body:
      - uid(String): 아이디 값인 이메일
      - pw(String): 사용자 비밀번호
    - response header: X
    - response body:
      - isSuccess(Boolean): 로그인 성공 여부
      - uname(String): 사용자 이름
  - `POST api/account/signup`: 회원정보를 받으면 사용자이름과 성공 여부를 응답합니다.
    - request header: X
    - params: X
    - request body:
      - uname(String): 사용자 이름
      - uid(String): 아이디 값인 이메일
      - pw(String): 사용자 비밀번호
    - response header: X
    - response body:
      - isSuccess(Boolean): signup 성공 여부
      - uname(String): 사용자 이름
  - `GET api/account/uname/{id}`: 아이디값을 쿼리스트링으로 받으면 사용자 이름을 응답합니다.
    - request header: X
    - params: id(String)
    - request body: X
    - response header: X
    - response body:
      - uname(String): 사용자 이름
  - `GET api/account/profile/{id}`: 아이디값을 쿼리스트링으로 받으면 프로필 이미지를 응답합니다.
    - request header: X
    - params: id(String)
    - request body: X
    - response header: X
    - response body:
      - profile(File): 프로필 이미지
  - `POST api/account/profile`: 유저정보와 프로필 이미지를 받으면 성공 여부를 응답합니다.
    - request header: X
    - params: X
    - request body:
      - profile(File): 프로필 이미지 파일
      - uname(String): 사용자 이름
    - response header: X
    - response body:
      - isSuccess(Boolean): 프로필 등록 성공 여부
- Image API
  - `POST api/img/upload`: 이미지 정보와 사용자 정보를 받으면 성공 여부를 응답합니다.
    - request header: X
    - params: X
    - request body:
      - img(File): 이미지 파일
      - imgName(String): 이미지 이름
      - uname(String): 업로드한 사용자의 이름
    - response header: X
    - response body:
      - isSuccess(Boolean): 업로드 성공 여부
  - `GET api/img/name`: 모든 이미지의 이름을 응답합니다.
    - request header: X
    - params: X
    - request body: X
    - response header: X
    - response body:
      - imgnames(String[]): 모든 이미지의 파일 이름