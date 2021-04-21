import React from "react";
import { Button, TextField } from "@material-ui/core";
import "../global.css";
import { Get } from "../_Utils/ApiCaller";
import { UserInfo } from "./UserInfo";

export class Select extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: "",
      user: null,
    };
  }

  selectByIdButtonHandler(id) {
    Get("http://localhost:5000/user/" + id).then((data) => {
      this.setState({
        user: data,
      });
    });
  }

  render() {
    return (
      <div className="AppContent">
        <h1>Select User</h1>
        <form>
          <TextField
            fullWidth
            type="number"
            id="user-id-input"
            label="User ID"
            onChange={(e) => this.setState({ id: e.target.value })}
            error={this.state.id === ""}
            value={this.state.id}
            helperText={this.state.id === "" ? "This cannot be empty!" : ""}
          />
          <Button
            variant="contained"
            color="primary"
            disabled={this.state.id === ""}
            onClick={() => this.selectByIdButtonHandler(this.state.id)}
          >
            Select By ID
          </Button>
        </form>
        {this.state.user !== null &&
          this.state.user !== undefined &&
          this.state.user.id !== null && (
            <UserInfo
              id={this.state.user.id}
              name={this.state.user.name}
              secretKey={this.state.user.secret_key}
            />
          )}
      </div>
    );
  }
}
