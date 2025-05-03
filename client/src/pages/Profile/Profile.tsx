import { useEffect } from "react";
import Navbar from "../../components/Navbar/Navbar";
import { useState } from "react";
import { FaUserAlt } from "react-icons/fa";
import GetProfile from "../../utils/GetProfile";
import { useNavigate, useParams } from "react-router";
import Cookies from "js-cookie";
import GetProfileByUserID from "../../utils/GetProfileByUserID";
import Error from "../../components/Error/Error";
import axios from "axios";
import Posts from "../../components/Posts/Posts";

const Profile = () => {
  const { userid } = useParams<{ userid: string }>();
  const navigate = useNavigate();
  let [userData, setUserData] = useState<any>(null);
  let [posts, setPosts] = useState<any>([]);
  let [error, setError] = useState<string | null>(null);
  function getUserPosts() {
    axios.get("/api/userposts/" + userid).then((res) => {
      setPosts(res.data);
      console.log(res.data);
    }).catch((err) => {
      console.log(err);
    });
  }
  function getProfileByUsedID() {
    GetProfileByUserID(Cookies.get("token"), userid)
      .then((res) => {
        console.log(res);
        setUserData(res);
      })
      .catch((err: string) => {
        setError(err);
        if (err === "Unauthorized") {
          navigate("/signin");
        }
      });
  }
  function getProfileByToken() {
    GetProfile(Cookies.get("token"))
      .then((res) => {
        console.log(res);
        setUserData(res);
      })
      .catch((err: string) => {
        console.log(err);
        if (err === "Unauthorized") {
          navigate("/signin");
        }
      });
  }
  useEffect(() => {
    if (userid) {
      setUserData(null);
      setPosts([]);
      getProfileByUsedID();
      getUserPosts();
    } else {
      getProfileByToken();
    }
  }, [userid]);
  useEffect(() => {
    if (userData && !userid) {
      if (userData.user.id) {
        navigate("/profile/" + userData.user.id);
      }
    }
  }, [userData]);
  return (
    <div className="w-full">
      <Navbar />
      {error ? (
        <Error message={error} />
      ) : (
        <div className="w-full flex flex-col items-center">
          {userData ? (
            <div className="text-white mt-10">
              <div className="flex justify-center items-center">
                <FaUserAlt className="text-9xl bg-blue-600 p-5 rounded-4xl" />
              </div>
              <h1 className="text-center text-white text-3xl">
                {userData.user.username}
              </h1>
              <h2 className="text-center text-white text-xl">
                {userData.user.email}
              </h2>
              <div className="flex justify-center items-center mt-2 gap-2">
                {
                  userData.isFollowing && !userData.isMe ? (
                    <button
                      className="bg-red-500 hover:cursor-pointer text-white px-4 py-2 rounded hover:bg-red-600"
                      onClick={() => {
                        axios
                          .get(`/api/follow/${userid}`, { headers: { Authorization: `${Cookies.get("token")}` } })
                          .then((res) => {
                            if (res.data.followers != undefined) {
                              setUserData((prevUserData: any) => ({
                                ...prevUserData,
                                user: {
                                  ...prevUserData.user,
                                  followers: res.data.followers,
                                },
                                isFollowing: res.data.isFollowing,
                              }));
                            }
                          })
                          .catch((err) => {
                            console.error("Error unfollowing user:", err);
                          });
                      }}
                    >
                      Unfollow
                    </button>
                  ) : !userData.isMe ? <button
                  className="bg-blue-500 hover:cursor-pointer text-white px-4 py-2 rounded hover:bg-blue-600"
                  onClick={() => {
                    axios
                    .get(`/api/follow/${userid}`, { headers: { Authorization: `${Cookies.get("token")}` } })
                    .then((res) => {
                      if (res.data.followers != undefined) {
                        setUserData((prevUserData: any) => ({
                          ...prevUserData,
                          user: {
                            ...prevUserData.user,
                            followers: res.data.followers,
                          },
                          isFollowing: res.data.isFollowing,
                        }));
                      }
                    })
                    .catch((err) => {
                      console.error("Error following user:", err);
                    });
                  }}
                >
                  Follow
                </button> : null

                }
              </div>
              <h2 className="text-xl text-center">{userData.user.followers} Followers / {userData.user.following} Following </h2>
            </div>
          ) : null}
          <h1 className="text-2xl text-white text-center">User posts</h1>
          {posts && userData ? (
            <Posts editmode={userData.isMe} posts={posts} />
          ) : null
          }
        </div>
      )}
    </div>
  );
};

export default Profile;
