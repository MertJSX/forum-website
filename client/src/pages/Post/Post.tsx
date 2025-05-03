import React, { useEffect } from "react";
import Navbar from "../../components/Navbar/Navbar";
import { useParams, Link } from "react-router";
import axios from "axios";
import { FaUserAlt } from "react-icons/fa";
import CreateNewComment from "../../components/CreateNewComment/CreateNewComment";
import moment from "moment";
import { BiSolidUpvote } from "react-icons/bi";
import Cookies from "js-cookie";

const Post: React.FC = () => {
  const { postid } = useParams<{ postid: string }>();
  const [post, setPost] = React.useState<any>(null);
  const [comments, setComments] = React.useState<any>([]);

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

  function getComments() {
    axios
      .get(`/api/posts/${postid}/comments`)
      .then((res) => {
        console.log("Comments:");
        console.log(res.data);
        setComments(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }

  useEffect(() => {
    getPost();
    getComments();
  }, []);

  return (
    <div>
      <Navbar />
      {post ? (
        <div className="w-3/4 bg-gray-700 flex flex-col mx-auto justify-center items-center mt-10 text-white p-4 rounded-lg">
          <div className="flex w-full justify-left ml-0 items-center mb-2 gap-2">
            <FaUserAlt className="text-6xl bg-blue-500 p-1 rounded-xl" />
            <Link to={`/profile/${post.userid}`} className="text-white hover:text-blue-200 text-2xl text-left">{post.author} <span className="text-gray-400 text-sm select-none italic">{ moment.unix(post.created_at).format("HH:mm D MMM YYYY") }</span></Link>
          </div>

          <h1 className="text-white text-3xl text-center w-full">{post.title}</h1>
          <p className="text-white text-lg text-justify">{post.content}</p>
          <div className="flex w-full justify-left items-center mt-4">
            <button
              className="flex items-center gap-2 bg-blue-600 hover:bg-blue-700 hover:cursor-pointer text-white font-bold py-2 px-4 rounded"
              onClick={() => {
                axios
                  .get(`/api/upvote/${postid}`, {
                    headers: { Authorization: Cookies.get("token") }})
                  .then((res) => {
                    console.log(res.data);
                    if (res.data.upvotes != undefined) {
                      setPost((prevPost: any) => ({
                        ...prevPost,
                        upvotes: res.data.upvotes,
                      }));
                    }
                  })
                  .catch((err) => {
                    console.log(err);
                  });
              }}
            >
              <BiSolidUpvote /> Upvote
            </button>
            {post.upvotes > 0 ? (
                <span className="ml-4 text-lg">{post.upvotes} Upvotes</span>
                ) : (
                <span className="ml-4 text-lg">No Votes Yet</span>
            )}
          </div>
        </div>
      ) : (
        <div className="flex justify-center items-center mt-10">
          <h1 className="text-white text-2xl">Loading...</h1>
        </div>
      )}
      <div className="flex flex-col justify-center items-center mt-10">
        <h1 className="text-white text-2xl">Comments</h1>
        <CreateNewComment getComments={getComments} postId={postid} />
        <div className="w-3/4 flex flex-col mx-auto justify-center items-center mt-4 text-white p-4 rounded-lg">
            {comments ? (
                comments.map((comment: any) => (
                    <div
                        key={comment.id}
                        className="w-full p-4 rounded-lg mb-4 flex flex-col"
                    >
                        <div className="flex items-center gap-2 mb-2">
                            <FaUserAlt className="text-6xl bg-blue-600 p-1 rounded-xl" />
                            <h2 className="text-white hover:text-blue-300 hover:cursor-pointer text-2xl flex justify-center items-center gap-2">
                                <Link to={`/profile/${comment.user_id}`}>{comment.author}</Link>
                                
                                <span className="text-gray-400 text-sm select-none italic">
                                    {moment.unix(comment.created_at).format("HH:mm D MMM YYYY")}
                                </span>
                            </h2>
                        </div>
                        <p className="text-white bg-gray-700 rounded-xl p-4 text-xl">{comment.comment}</p>
                    </div>
                ))
            ) : (
                <div className="text-gray-400 text-lg italic">No comments yet. Be the first to comment!</div>
            )}
        </div>
      </div>
    </div>
  );
};

export default Post;
