type ErrorProps = {
    message: string
    code?: number
}

const Error = ({message, code}:ErrorProps) => {
  return (
    <div className="flex flex-col items-center justify-center mx-auto mt-20 w-1/2 h-64 rounded-2xl bg-gray-700 text-white">
      <h1 className="text-4xl font-bold mb-4 text-blue-300">Error {code}</h1>
        <p className="text-lg">{message}</p>
    </div>
  )
}

export default Error