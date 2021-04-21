import React from "react";

export class UserInfo extends React.Component {
  render() {
    return (
      <div>
        <h3>{this.props.name}</h3>
        <p>ID: {this.props.id}</p>
        <p>Secret Key: {this.props.secretKey}</p>
      </div>
    );
  }
}
