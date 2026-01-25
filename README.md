POST   /auth/register
POST   /auth/login

GET    /users          (admin only)
GET    /users/:id      (admin or own profile)
DELETE /users/:id      (admin only)

GET    /tasks          (admin: all, user: own)
POST   /tasks
GET    /tasks/:id
PUT    /tasks/:id
DELETE /tasks/:id