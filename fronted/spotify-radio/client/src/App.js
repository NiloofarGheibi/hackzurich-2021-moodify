import './App.css';
import "bootstrap/dist/css/bootstrap.min.css";
import AdminView from "./AdminView";
import Dashboard from "./Dashboard";
import Login from "./Login";
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';

const code = new URLSearchParams(window.location.search).get('code')

function App() {
 return(
 <Router>
  <div>
   <Switch>
    <Route exact path="/" component={() => code ? <Dashboard code={code} /> : <Login />} />
    <Route exact path="/admin" component={() => <AdminView />} />
   </Switch>
  </div>
 </Router>
 )

}

export default App;
