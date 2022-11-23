// import { Box } from "@mantine/core"
import "./App.css";
import useSWR from "swr";
import AddTodo from "./components/AddTodo";

export interface Todo{
  id: number;
  title: string;
  body: string;
  done: boolean
}

export const ENDPOINT = "http://localhost:4000"

const fetcher = (url:string) => fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

function App() {
 

  const {data, mutate} = useSWR<Todo[]>("api/todos", fetcher);

  return (
    <div >
      {JSON.stringify(data)}
      <AddTodo mutate={mutate}/>
    </div>
  )
}

export default App
