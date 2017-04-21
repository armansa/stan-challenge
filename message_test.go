package main

import (
  "strconv"
  "testing"
)

var message = Message {
  Payload: []Show {
    Show {
      Drm: true,
      EpisodeCount: 3,
      Img: Image { "http://catchup.com.au/1280.jpg" },
      Slug: "show/16kidsandcounting",
      Title: "16 Kids and Counting",
    },
    Show {
      Slug: "show/seapatrol",
      Title: "Sea Patrol",
    },
    Show {
      Drm: true,
      EpisodeCount: 2,
      Img: Image { "http://catchup.com.au/1280.jpg" },
      Slug: "show/thetaste",
      Title: "The Taste (Le GoÃ»t)",
    },
  },
  Skip: 0,
  Take: 10,
  TotalRecords: 75,
}

func TestFilter(t *testing.T) {
  response := message.Filter()
  if len(response) != 2 {
    t.Error("Expected response length to be 2 but is ", strconv.Itoa(len(response)))
  }
}
