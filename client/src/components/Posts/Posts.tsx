import { useNavigate } from "react-router";
import TextTruncate from "../../utils/TextTruncate";
import "./Posts.css";
import axios from "axios";
import Cookies from "js-cookie";
import moment from "moment";

type PostsProps = {
  posts?: any[];
  editmode?: boolean;
};

const Posts = ({ posts, editmode }: PostsProps) => {
  const navigate = useNavigate();

  function handlePostClick(postId: string) {
    navigate(`/post/${postId}`);
  }

  function handleProfileClick(userId: string) {
    navigate(`/profile/${userId}`);
  }

  return (
    <div className="w-full">
      {posts ? (
        <div className="posts flex flex-col items-center mt-10 w-1/2 justify-center mx-auto">
          {posts.map((post, index) => (
            <div
              key={index}
              className="post bg-gray-700 hover:bg-gray-600 transition-all text-white p-4 rounded-lg mb-4 w-full"
            >
              <div className="flex items-center gap-2">
                <h2
                  onClick={() => {
                    handlePostClick(post.id);
                  }}
                  className="text-xl font-bold hover:text-blue-300 cursor-pointer"
                >
                  {post.title}
                </h2>
                -
                <h2
                  onClick={() => {
                    handleProfileClick(post.userid);
                  }}
                  className="text-gray-400 hover:text-blue-300 text-lg font-bold italic cursor-pointer"
                >
                  {post.author}
                </h2>
              </div>
              <p>
                <TextTruncate text={post.content} />
              </p>
                <div className="flex items-center gap-2 mt-2">
                    <span className="text-gray-400 text-sm select-none italic">
                    {moment.unix(post.created_at).format("HH:mm D MMM YYYY")}
                    </span>
                    <span className="text-gray-400 text-sm select-none italic">
                    {post.upvotes} Upvotes
                    </span>
                </div>
              {editmode ? (
                <div className="flex justify-start items-center mt-2 gap-2">
                  <button
                    className="bg-blue-500 hover:bg-blue-600 cursor-pointer text-white px-4 py-2 rounded"
                    onClick={() => {
                      navigate(`/editpost/${post.id}`);
                    }}
                  >
                    Edit
                  </button>
                  <button
                    className="bg-red-500 hover:bg-red-600 cursor-pointer text-white px-4 py-2 rounded"
                    onClick={() => {
                      axios
                        .delete(`/api/posts/${post.id}`, {
                          headers: { Authorization: Cookies.get("token") },
                        })
                        .then((res) => {
                          console.log(res.data);
                          window.location.reload();
                        })
                        .catch(() => {
                          alert("Error deleting post");
                        });
                    }}
                  >
                    Delete
                  </button>
                </div>
              ) : null}
            </div>
          ))}
        </div>
      ) : (
        <h1 className="text-white text-2xl text-center">No posts available</h1>
      )}
    </div>
  );
};

export default Posts;
