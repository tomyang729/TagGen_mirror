package main


import (
  "encoding/json"
  "sort"
)

type ClarifaiObj struct {
  Outputs []ClarifaiData `json:"outputs"`
}

type ClarifaiData struct {
  Data TagsObj `json:"data"`
}

type ClarifyTag struct {
  Name  string  `json:"name"`
  Value float64 `json:"value"`
}

type TagsObj struct {
  Tags []ClarifyTag `json:"concepts"`
}

func filterImageTags(tags []ClarifyTag) []ClarifyTag {
  res := []ClarifyTag{}

  score := func(t1, t2 *ClarifyTag) bool {
    return t1.Value > t2.Value
  }

  for _, tag := range tags {
    if tag.Value >= 0.72 {
      res = append(res, tag)
      // fmt.Println(tag.Name)
    }
  }

  By(score).Sort(res)
  if len(res) > 10 {
    res = res[:10]
  }

  fmt.Println(res)
  return res
}

//custom sorter
type tagsSorter struct {
  tags []ClarifyTag
  by   func(p1, p2 *ClarifyTag) bool // Closure used in the Less method.
}

type By func(p1, p2 *ClarifyTag) bool

func (by By) Sort(tags []ClarifyTag) {
  ps := &tagsSorter{
    tags: tags,
    by:   by, // The Sort method's receiver is the function (closure) that defines the sort order.
  }
  sort.Sort(ps)
}

func (s *tagsSorter) Len() int {
  return len(s.tags)
}

func (s *tagsSorter) Swap(i, j int) {
  s.tags[i], s.tags[j] = s.tags[j], s.tags[i]
}

func (s *tagsSorter) Less(i, j int) bool {
  return s.by(&s.tags[i], &s.tags[j])
}
