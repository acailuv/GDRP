import React from "react";
import "../global.css";

export class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <div className="AppContent">
        <p>This is a demo project for this template stack.</p>
        <p>GDRP: Go-Docker-React-PostgreSQL</p>
      </div>
    );
  }
}
