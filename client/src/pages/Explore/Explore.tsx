import Navbar from "../../components/Navbar/Navbar"
import CreateNewPost from "../../components/CreateNewPost/CreateNewPost"

const Explore = () => {
  return (
    <div>
        <Navbar />
        <CreateNewPost />
        <h1 className="text-white text-2xl text-center">Explore newest posts!</h1>
    </div>
  )
}

export default Explore