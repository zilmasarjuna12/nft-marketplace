mutation createItem {
  createItem(
    creatorId:"9bbbd26e-4738-4539-8ab4-e91f085231f5",
    input: {
      name: "tes",
      rating: 100,
      category:"animation",
      image:"http:zilmas.com",
      reputation: 200,
      price: 1000,
      availibility: 10
    }
  ) {
    id
  }
}