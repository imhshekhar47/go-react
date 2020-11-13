import axios from 'axios';
import React, { useEffect, useState } from 'react';
import './App.scss';
import gopher from './images/gopher.png';

interface IListItemProps {
  name: string
  value?: any
};
const ListItem: React.FC<IListItemProps> = (props) => {
  return (
    <div className="list-item">
      <label className="item-name">{props.name}</label>
      <label className="item-value">{props.value}</label>
    </div>
  )
}

interface IListProps {
  title: string,
  children: React.ReactNode | null
};
const List: React.FC<IListProps> = (props) => {
  return (
    <div className="list-container">
      <label className="list-title">{props.title}</label>
      {props.children}
    </div>
  )
}

const ObjectView = (props: {title: string, data: Object}) => {
  return(
    <List title={props.title}>
      {
        Object.entries(props.data).map(([key, value]) => (
          <ListItem key={key} name={key} value={value} />
        ))
      }
    </List>
  )
}

const App = () => {

  const [activeId, setActiveId] = useState<string|null>(null);
  const [appInfo, setAppInfo] = useState<Object|null>(null);
  const [hostInfo, setHostInfo] = useState<Object|null>(null);
  const [podInfo, setPodInfo] = useState<Object|null>(null);


  useEffect(() => {
    console.log("On First load");
    loadAppInfo();
  }, [])

  const loadAppInfo = () => {
    setActiveId("app");
    axios.get("/api")
      .then(response => {
        const {data} = response;
        console.log(data);
        setAppInfo(data);
      })
      .catch(err => {
        console.log(err);
      })
  }

  const loadHostInfo = () => {
    setActiveId("host");
    axios.get("/api/system/host")
      .then(response => {
        const {data} = response;
        console.log(data);
        setHostInfo(data);
      })
      .catch(err => {
        console.log(err);
      })
  }

  const loadPodInfo = () => {
    setActiveId("pod");
    axios.get("/api/system/pod")
      .then(response => {
        const {data} = response;
        console.log(data);
        setPodInfo(data);
      })
      .catch(err => {
        console.log(err);
      })
  }

  return (
    <div className="main-container">
      <header className="main-header contained">
        <div className="box avatar-box">
          <img src={gopher} className="avatar" alt="logo" />
        </div>
        <div className="box content-box">
          <h1 className="main-heading">Hi There,</h1>
          <h5 className="main-sub-heading">Looks like you are early.</h5>
          <div className="description">
            <p>This is a placeholde app. The work is in progress and this will soon be updated.</p>
            <p>Please come back in few days.</p>
          </div>
        </div>
      </header>
      <div className="system-info contained">
        <div className="btn-group">
          <span className={ activeId == "app" ? "btn active" : "btn"} onClick={loadAppInfo}>App Info</span>
          <span className={ activeId == "host" ? "btn active" : "btn"} onClick={loadHostInfo}>Host Info</span>
          <span className={ activeId == "pod" ? "btn active" : "btn"} onClick={loadPodInfo}>Pod Info</span>
        </div>
        { appInfo && activeId === "app" && <ObjectView title="App Information" data={appInfo} />}
        { hostInfo && activeId === "host" && <ObjectView title="Host Information" data={hostInfo} />}
        { podInfo && activeId === "pod" && <ObjectView title="Pod Information" data={podInfo} />}
      </div>
      <div className="stretcher" />
      <div className="developer-info">

        <div className="contained">
          <small>Developed by <span className="keyword">@imhshekhar47</span></small>
        </div>
      </div>
    </div>
  );
}

export default App;
