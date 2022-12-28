import './App.css';
import {BrowserRouter as Router} from "react-router-dom";
import { HomePageLinks, RoutingPages } from './pages/Routing';

function App() {
  return (
    <div className="App">
      <Router>
        <HomePageLinks/>
        <RoutingPages/>
      </Router>
    </div>
  );
}

export default App;
