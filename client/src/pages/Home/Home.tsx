import Navbar from "../../components/Navbar/Navbar"
import CreateNewPost from "../../components/CreateNewPost/CreateNewPost"
import Cookies from "js-cookie"
import { useEffect, useState } from "react"
import axios from "axios"
import Posts from "../../components/Posts/Posts"


const Home = () => {
  const [posts, setPosts] = useState<any>([]);

  function getPosts() {
    axios.get("/api/followed-users-posts", {
        headers: { Authorization: Cookies.get("token") },})
      .then((res) => {
        console.log(res.data);
        setPosts(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }


  useEffect(() => {
    getPosts();
  }, []);
  return (
    <div>
        <Navbar />
        <h1 className='text-2xl text-center text-sky-400 font-bold'>Welcome to the Forum website!</h1>
        {
          Cookies.get("token") ? (
            <div>
              <CreateNewPost />
              <h1 className='text-xl text-center text-sky-200 my-10'>Here are posts of users that you follow!</h1>
              <Posts posts={posts} />
            </div>
          ) : (
            <h1 className='text-2xl text-center text-sky-400 font-bold'>Login to create a new post!</h1>
          )
        }
    </div>
  )
}

export default Home