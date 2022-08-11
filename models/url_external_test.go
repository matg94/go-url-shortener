package models

import (
	"testing"
)

func TestShortenRequestFromJson(t *testing.T) {
	json_data := []byte("{\"url\": \"http://google.com\"}")
	url, err := ShortenRequestFromJson(json_data)

	if err != nil {
		t.Log("expected no errors but got", err)
		t.Fail()
	}

	expectedURL := URLShortenRequest{URL: "http://google.com"}

	if url != expectedURL {
		t.Log("expected url object to contain http://google.com, but got", url)
		t.Fail()
	}
}

func TestShortenRequestFromJsonMissingURL(t *testing.T) {
	json_data := []byte("{\"test\": \"http://google.com\"}")
	url, err := ShortenRequestFromJson(json_data)

	if err != ErrUrlNotDefined {
		t.Log("expected ErrUrlNotDefined but got", err)
		t.Fail()
	}

	expectedURL := URLShortenRequest{}

	if url != expectedURL {
		t.Log("expected url object be empty but got", url)
		t.Fail()
	}
}

func TestShortenRequestFromJsonParsingError(t *testing.T) {
	json_data := []byte("{\"http://google.com\"}")
	url, err := ShortenRequestFromJson(json_data)

	if err == nil || err == ErrUrlNotDefined {
		t.Log("expected some parsing error but got", err)
		t.Fail()
	}

	expectedURL := URLShortenRequest{}

	if url != expectedURL {
		t.Log("expected url object be empty but got", url)
		t.Fail()
	}
}

func TestLongRequestFromJsonJson(t *testing.T) {
	json_data := []byte("{\"url\": \"testhash\"}")
	hash, err := LongRequestFromJson(json_data)

	if err != nil {
		t.Log("expected no errors but got", err)
		t.Fail()
	}

	expectedHash := URLElongateResponse{Hash: "testhash"}

	if hash != expectedHash {
		t.Log("expected url object to contain testhash, but got", hash)
		t.Fail()
	}
}

func TestLongRequestFromJsonMissingURL(t *testing.T) {
	json_data := []byte("{\"test\": \"testhash\"}")
	hash, err := LongRequestFromJson(json_data)

	if err != ErrUrlNotDefined {
		t.Log("expected ErrUrlNotDefined but got", err)
		t.Fail()
	}

	expectedHash := URLElongateResponse{}

	if hash != expectedHash {
		t.Log("expected url object be empty but got", hash)
		t.Fail()
	}
}

func TestLongRequestFromJsonParsingError(t *testing.T) {
	json_data := []byte("{\"testHash\"}")
	hash, err := LongRequestFromJson(json_data)

	if err == nil || err == ErrUrlNotDefined {
		t.Log("expected some parsing error but got", err)
		t.Fail()
	}

	expectedHash := URLElongateResponse{}

	if hash != expectedHash {
		t.Log("expected url object be empty but got", hash)
		t.Fail()
	}
}
