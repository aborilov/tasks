## Task
Develop a task management system. Each task should have its own status, description, and users assigned to it. 
Statuses:  new, in progress, completed, archived.
The system should provide the user roles. Only the administrator should be able to archive tasks.
Users assigned to tasks receive notifications about changes by email.
Users should be able to comment on the tasks. Only administrators can delete comments.

## Comments
Here I created only business layer because I have no information about what repository should be used(mysql, postgresq, etc), same with API, do we need REST or not. API and Repository implementation can be added easily, the business logic layer will stay untouched.
I really don't like to have `User` in context and checking permission in the business layer, I usually like to do this in the API layer, but because there is no API implemented within this task, I have to put it to the business layer.
