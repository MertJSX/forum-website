import axios from "axios";
import { useState } from "react";
import Cookies from "js-cookie";
import { useNavigate } from "react-router";

type CreateNewPostProps = {
  getPosts?: () => void;
};

const CreateNewPost = ({getPosts}: CreateNewPostProps) => {
  const navigate = useNavigate();
  let [title, setTitle] = useState<string>("");
  let [content, setContent] = useState<string>("");
  function handleCreateNewPost(e: React.FormEvent) {
    e.preventDefault();
    axios
      .post(
        "/api/create-post",
        {
          title: title,
          content: content,
        },
        {
          headers: { Authorization: Cookies.get("token") },
        }
      )
      .then((res) => {
        navigate("/post/" + res.data.postID);
        setTitle("");
        setContent("");
        if (getPosts) {
          getPosts();
        }
      })
      .catch((err) => {
        console.log(err);
        if (err.response.status === 401) {
          console.log("Unauthorized");
        } else if (err.response.status === 400) {
          console.log("Bad Request");
        } else {
          console.log("Unknown error");
        }
      });
  }
  return (
    <div>
      <h1 className="text-white text-2xl text-center">Create a new Post!</h1>
      <form
        onSubmit={handleCreateNewPost}
        className="flex flex-col items-center mt-10 w-1/2 justify-center mx-auto"
      >
        <input
          type="text"
          placeholder="Title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          className="bg-gray-700 text-white p-2 rounded-lg mb-4 w-full"
        />
        <textarea
          placeholder="Content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          className="bg-gray-700 text-white p-2 rounded-lg mb-4 w-full h-32"
        ></textarea>
        <button
          type="submit"
          className="bg-blue-600 hover:bg-blue-500 hover:cursor-pointer w-full text-white p-2 rounded-lg"
        >
          Create Post
        </button>
      </form>
    </div>
  );
};

export default CreateNewPost;
