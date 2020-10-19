import React from 'react';
import {BrowserRouter as Router,Link,Switch,Route} from 'react-router-dom';
import './NavComponent.css';

function NavComponent(props)
{
  
    return (
          <React.Fragment>
          <div className="nav-wrapper">
          <nav className="nav-bar">
            <li className="nav-item green"><Link className="nav-item" to="/">Home</Link></li>
            <li className="nav-item"><Link className="nav-item" to="/concept">Concept</Link></li>
            <li className="nav-item"><Link className="nav-item " to="/fuzzers">Fuzzers</Link></li>
            <li className="nav-item"><Link className="nav-item " to="/authors">Authors</Link></li>
            <li className="nav-item"><Link className="nav-item" to="/donate">Donate</Link></li>
          </nav>
          

          </div>
          </React.Fragment>
         
     
    )
}

export default NavComponent;
