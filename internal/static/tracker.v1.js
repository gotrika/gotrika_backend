function dec2hex(dec) {
  return dec.toString(16).padStart(2, "0");
}

function generateId(len) {
  var arr = new Uint8Array((len || 40) / 2);
  window.crypto.getRandomValues(arr);
  return Array.from(arr, dec2hex).join("");
}

function getNow() {
  var now = new Date();
  return now;
}

function getUTC(now) {
  var utc_timestamp = Date.UTC(
    now.getUTCFullYear(),
    now.getUTCMonth(),
    now.getUTCDate(),
    now.getUTCHours(),
    now.getUTCMinutes(),
    now.getUTCSeconds(),
    now.getUTCMilliseconds()
  );
  return Math.floor(utc_timestamp / 1000);
}

var GOtrika = {
  code: null,
  url: null,
  heartbeatTaskId: null,
  heartbeatID: generateId(25),
  validateEvent: function (e) {
    console.log(e.target);
    var valid_targets = ["button", "form", "a"];
    if (valid_targets.includes(e.target.localName)) {
      GOtrika._sendEvent(e);
    }
  },
  _sendEvent: function (e) {
    var hit_url = e.target.href || "";
    var tracked_event_data = {
      class_name: e.target.className,
      event_type: e.target.localName,
      page: window.location.href,
      page_title: document.title,
      visitor_id: GOtrika.getVisitorId(),
      session_id: GOtrika.getSessionId(),
      client_timestamp: Math.floor(getNow().getTime() / 1000),
      hit_url: hit_url,
    };
    var data = {
      site_id: GOtrika.code,
      type: "event",
      hash_id: generateId(),
      timestamp: getUTC(getNow()),
      tracker_data: tracked_event_data,
    };
    try {
      var xhr = new XMLHttpRequest();
      xhr.open(
        "POST",
        GOtrika.url,
        true
      );
      xhr.setRequestHeader("Content-Type", "application/json");
      xhr.send(JSON.stringify(data));
    } catch (e) {
      console.log(e);
    }
  },
  getHashId: function () {
    var hash_id = localStorage.getItem("_GOtrika_hash_id") || generateId();
    localStorage.setItem("_GOtrika_hash_id", hash_id);
    return hash_id;
  },
  Init: function (code, url) {
    GOtrika.code = code;
    GOtrika.url = url;
  },
  getVisitorId: function () {
    var visitor_id =
      localStorage.getItem("_GOtrika_visitor_id") || generateId(25);
    localStorage.setItem("_GOtrika_visitor_id", visitor_id);
    return visitor_id;
  },
  getSessionId: function () {
    var session_id = localStorage.getItem("_GOtrika_session_id");
    return session_id;
  },
  getSessionTimestamp: function () {
    var session_timestamp = localStorage.getItem("_GOtrika_session_timestamp");
    return session_timestamp;
  },
  setVisitId: function () {
    var session_id = localStorage.getItem("_GOtrika_session_id") || generateId();
    var now = getNow();
    var utc_now = getUTC(now);
    var session_timestamp =
      localStorage.getItem("_GOtrika_session_timestamp") || 0;
    if (session_timestamp === 0) {
      localStorage.setItem("_GOtrika_session_timestamp", utc_now);
    } else if (session_timestamp + 1800 < utc_now) {
      session_id = generateId();
      hash_id = generateId();
      localStorage.setItem("_GOtrika_hash_id", hash_id);
      localStorage.setItem("_GOtrika_session_timestamp", utc_now);
    }
    localStorage.setItem("_GOtrika_session_id", session_id);
  },
  _sendVisit: function () {
    var tracker_data = {
      referrer: document.referrer,
      location: window.location.href,
      user_screen_width: screen.width,
      user_screen_height: screen.height,
      client_timezone: Intl.DateTimeFormat().resolvedOptions().timeZone,
      user_agent: navigator.userAgent,
      visitor_id: GOtrika.getVisitorId(),
      session_id: GOtrika.getSessionId(),
      session_timestamp: GOtrika.getSessionTimestamp(),
    };
    var data = {
      site_id: GOtrika.code,
      type: "session",
      hash_id: generateId(),
      timestamp: getUTC(getNow()),
      tracker_data: tracker_data,
    };
    try {
      var xhr = new XMLHttpRequest();
      xhr.open(
        "POST",
        GOtrika.url + "visit/" + "?site_id=" + GOtrika.code,
        true
      );
      xhr.setRequestHeader("Content-Type", "application/json");
      xhr.send(JSON.stringify(data));
    } catch (e) {
      console.log(e);
    }
  },
  newPageLoad: function (e) {
    console.log(e);
    if (GOtrika.heartbeatTaskId != null) {
      clearInterval(GOtrika.heartbeatTaskId);
    }
    GOtrika.setVisitId();
    GOtrika.heartbeatTaskId = setInterval(
      GOtrika._sendVisit,
      parseInt("900000000")
    );
    GOtrika._sendVisit();
  },
};
window.addEventListener("load", GOtrika.newPageLoad);
document.addEventListener("click", GOtrika.validateEvent);
document.addEventListener("submit", GOtrika.validateEvent);
