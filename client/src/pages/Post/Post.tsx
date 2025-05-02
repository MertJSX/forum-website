import React, { useEffect } from "react";
import Navbar from "../../components/Navbar/Navbar";
import { useParams } from "react-router";
import axios from "axios";
import { FaUserAlt } from "react-icons/fa";
import CreateNewComment from "../../components/CreateNewComment/CreateNewComment";

const Post: React.FC = () => {
  const { postid } = useParams<{ postid: string }>();
  const [post, setPost] = React.useState<any>(null);

  function getPost() {
    axios
      .get(`/api/posts/${postid}`)
      .then((res) => {
        console.log(res.data);
        setPost(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }

  useEffect(() => {
    getPost();
    
  }, []);

  return (
    <div>
      <Navbar />
      {post ? (
        <div className="w-3/4 bg-gray-800 flex flex-col mx-auto justify-center items-center mt-10 text-white p-4 rounded-lg">
          <div className="flex w-full justify-left ml-0 items-center mb-2 gap-2">
            <FaUserAlt className="text-6xl bg-neutral-600 p-1 rounded-xl" />
            <h1 className="text-white text-2xl text-left">{post.author}</h1>
          </div>

          <h1 className="text-white text-4xl text-left w-full">{post.title}</h1>
          <p className="text-white text-lg text-justify">{post.content}</p>
        </div>
      ) : (
        <div className="flex justify-center items-center mt-10">
          <h1 className="text-white text-2xl">Loading...</h1>
        </div>
      )}
      <div className="flex flex-col justify-center items-center mt-10">
        <h1 className="text-white text-2xl">Comments</h1>
        <CreateNewComment postId={postid} />
      </div>
    </div>
  );
};

export default Post;
