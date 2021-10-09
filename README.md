# APPOINTY-TASK

As port 1000 was free I decided to use that , 
These are the following commands for all the required functionalities:

- CREATE A NEW USER http://localhost:10000/users
<img src="https://github.com/sahilsaleem2907/appointy-task-api-instagram/blob/master/screenshots/create_user.PNG" width="600" />
- CREATE A NEW POST http://localhost:10000/posts
- <img src="https://github.com/sahilsaleem2907/appointy-task-api-instagram/blob/master/screenshots/create-post.PNG" width="600" />
- GET USER USING ID http://localhost:10000/users/sahilsal
- <img src="https://github.com/sahilsaleem2907/appointy-task-api-instagram/blob/master/screenshots/search_user_by_id.PNG" width="600" />
- GET ONE POST USING ID http://localhost:10000/posts/sahilsal
- <img src="https://github.com/sahilsaleem2907/appointy-task-api-instagram/blob/master/screenshots/get_post_by_id_unq.PNG" width="600" />
- GET ALL POSTS USING ID http://localhost:10000/posts/users/sahilsal
- <img src="https://github.com/sahilsaleem2907/appointy-task-api-instagram/blob/master/screenshots/get_post_by_id_ALL.PNG" width="600" />


By Default I have used 'sahilsal' as the id and made it global , as I have not taken any custom input from the user.

Backend architecturee:
- user Collection ![Image6](screenshots/user-db.PNG)
- post Collection ![Image7](screenshots/post-db.PNG)

- EncodingURL-compatible base64 format for password to be more secure 
- Pagination implemented

