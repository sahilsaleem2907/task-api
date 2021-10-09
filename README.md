# APPOINTY-TASK

As port 1000 was free I decided to use that , 
These are the following commands for all the required functionalities:

- CREATE A NEW USER http://localhost:10000/users
![Image1](screenshots/create_user.PNG) 
- CREATE A NEW POST http://localhost:10000/posts
![Image2](screenshots/create-post.PNG) 
- GET USER USING ID http://localhost:10000/users/sahilsal
![Image3](screenshots/search_user_by_id.PNG)
- GET ONE POST USING ID http://localhost:10000/posts/sahilsal
![Image4](screenshots/get_post_by_id_unq.PNG)
- GET ALL POSTS USING ID http://localhost:10000/posts/users/sahilsal
![Image5](screenshots/get_post_by_id_ALL.PNG)

By Default I have used 'sahilsal' as the id and made it global , as I have not taken any custom input from the user.

Backend architecturee:
- user Collection ![Image6](screenshots/user-db.PNG)
- post Collection ![Image7](screenshots/post-db.PNG)

- EncodingURL-compatible base64 format for password to be more secure 
- Pagination implemented

