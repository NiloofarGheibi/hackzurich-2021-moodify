import './App.css';
import "bootstrap/dist/css/bootstrap.min.css";
import AdminView from "./AdminView";
import Dashboard from "./Dashboard";
import Login from "./Login";

const code = new URLSearchParams(window.location.search).get('code')
function App() {
  //client
  return code ? <Dashboard code={code} /> : <Login />;

  //admin
 // return <AdminView />;
}

export default App;
