import axios from "axios";
import { useEffect, useState } from "react";
import Cookies from "js-cookie";
import { FaUserAlt } from "react-icons/fa";
import GetProfile from "../../utils/GetProfile";
import { useNavigate } from "react-router";

type CreateNewCommentProps = {
  getComments?: () => void;
  postId?: string;
};

const CreateNewComment = ({ getComments, postId }: CreateNewCommentProps) => {
  let [profile, setProfile] = useState<any>(null);
  let [comment, setComment] = useState<string>("");
  const navigate = useNavigate();
  function handleCreateNewPost(e: React.FormEvent) {
    e.preventDefault();
    axios
      .post(
        "/api/create-comment",
        {
          comment: comment,
          post_id: postId,
        },
        {
          headers: { Authorization: Cookies.get("token") },
        }
      )
      .then((res) => {
        console.log(res.data);
        setComment("");
        if (getComments) {
          getComments();
        }
      })
      .catch((err) => {
        if (err.response.status === 401) {
          navigate("/signin");
        }
      });
  }

  useEffect(() => {
    GetProfile(Cookies.get("token")).then((res) => {
        setProfile(res);
      }).catch((err: string) => {
        if (err === "Unauthorized") {
          navigate("/signin");
        }
      });
   }, []);
  return (
    <div className="w-full">
      <form
        onSubmit={handleCreateNewPost}
        className="flex flex-col items-center mt-10 w-2/3 justify-center mx-auto"
      >
        {
            profile ? (
                <div className="flex w-full justify-left ml-0 items-center mb-2 gap-2">
                    <FaUserAlt className="text-6xl bg-blue-600 p-1 rounded-xl text-white" />
                    <h1 className="text-white text-2xl text-left">{profile.user.username}</h1>
                </div>
                ) : null
        }
        <textarea
          placeholder="Your comment..."
          rows={4}
          value={comment}
          onChange={(e) => setComment(e.target.value)}
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

export default CreateNewComment;
