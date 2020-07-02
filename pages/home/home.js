Page({
  data: {
    showDialog: false,
    subscribeButton: {
      text: '预定',
      className: 'subscription-button',
      loading: false,
      disabled: false
    },
    games: [],
    isLoading: false,
    selectedGameIndex: 0,
    showFilterDialog: false,
    isSearching: false,
    filters: [{
        text: "所有区",
        value: "zone:all"
      },
      {
        text: "澳服",
        value: "zone:au"
      },
      {
        text: "美服",
        value: "zone:us"
      },
      {
        text: "收藏套装",
        value: "edition"
      },
      {
        text: "评分排序",
        value: "order:rating"
      },
      {
        text: "价格升序",
        value: "order:price-asc"
      },
      {
        text: "价格降序",
        value: "order:price-desc"
      },
      {
        text: "我的预定",
        value: "mine:subscription"
      }
    ],
    selectedFilterIndex: 0,
    inputShowed: false,
    inputVal: "",
    lastSearchedAt: 0
  },
  onLoad: function({
    gameName,
    featured
  }) {
    if (getApp().globalData.openid == '') {
      wx.redirectTo({
        url: '../launch/launch'
      })
    }
    this.setData({
      inputVal: gameName ? gameName.replace(/\+/g, " ") : '',
      selectedFilterIndex: featured ? 4 : 0
    }, () => this.fetchData(this))
  },
  fetchData: function(homePage) {
    homePage.setData({
      isLoading: true
    }, () => wx.request({
      method: 'GET',
      url: 'https://dw.dongwm.com/api/games',
      header: {
        "x-corran-token": getApp().globalData.openid
      },
      data: {
        "q": this.data.inputVal,
        "filter": this.data.filters[this.data.selectedFilterIndex].value,
      },
      success: function({
        data
      }) {
        const games = data['games']
        homePage.setData({
          games: games, // homePage.data.games.concat(games),
          isLoading: false
        })
      }
    }))
  },
  onPullDownRefresh: function() {
    this.reFetchData(this, () => {
      wx.stopPullDownRefresh()
    }, false)
  },
  reFetchData: function(homePage, callback, isLoading = true) {
    homePage.setData({
      isLoading: isLoading
    }, () => wx.request({
      method: 'GET',
      url: 'https://dw.dongwm.com/api/games',
      header: {
        "x-corran-token": getApp().globalData.openid
      },
      data: {
        "q": this.data.inputVal,
        "filter": this.data.filters[this.data.selectedFilterIndex].value,
      },
      success: function({
        data
      }) {
        const games = data['games']
        homePage.setData({
          games: games,
          isLoading: false
        }, callback)
      }
    }))
  },
  onReachBottom: function() {
    this.fetchData(this)
  },
  showFilter: function() {
    this.setData({
      showFilterDialog: true
    })
  },
  filterGames: function({
    detail
  }) {
    this.setData({
      showFilterDialog: false,
      selectedFilterIndex: detail.index,
      games: []
    }, () => this.fetchData(this))
  },
  clearInput: function() {
    this.setData({
      inputVal: ""
    }, () => this.reFetchData(this));
  },
  inputTyping: function(e) {
    this.setData({
      inputVal: e.detail.value
    }, () => this.search(this))
  },
  toDouban: function(e) {
    const game = this.data.games[this.data.selectedGameIndex];
    wx.navigateToMiniProgram({
      appId: 'wx2f9b06c1de1ccfca',
      path: 'page/subject/subject?type=game&id=' + game.douban_id,
      success(res) {
        // 打开成功
      }
    })
  },
  search: function() {
    this.setData({
      isSearching: true
    }, () => this.reFetchData(this, () => {
      this.setData({
        isSearching: false
      })
    }))
  },
  selectGame: function({
    currentTarget
  }) {
    const game = this.data.games[currentTarget.dataset.gameIndex];
    this.setData({
      showDialog: true,
      selectedGameIndex: currentTarget.dataset.gameIndex,
      subscribeButton: {
        text: game.subscribed ? '取消预定' : '确认预定',
        className: game.subscribed ? 'subscribed-button' : 'subscription-button',
        loading: false,
        disabled: false
      }
    })
  },
  unsubscribe: function() {
    const homePage = this;
    const game = this.data.games[this.data.selectedGameIndex];
    homePage.setData({
      subscribeButton: {
        ...homePage.data.subscribeButton,
        text: '',
        loading: true,
        disabled: true
      }
    })
    wx.request({
      method: 'POST',
      url: `https://dw.dongwm.com/api/game/${game.id}/unsubscribe`,
      header: {
        "x-corran-token": getApp().globalData.openid
      },
      success: function ({
        statusCode
      }) {
        if (statusCode == 200) {
          homePage.setData({
            subscribeButton: {
              text: '确认预定',
              className: 'subscription-button',
              loading: false,
              disabled: false
            },
            games: homePage.data.games.map(g => {
              if (g.id == game.id) {
                return {
                  ...game,
                  subscribed: false
                }
              } else {
                return g
              }
            })
          })
        }
      }
    })
  },
  subscribe: function() {
    const homePage = this;
    const game = this.data.games[this.data.selectedGameIndex];
    if (game.subscribed) {
      this.unsubscribe();
      return
    }
    homePage.setData({
      subscribeButton: {
        ...homePage.data.subscribeButton,
        text: '',
        loading: true,
        disabled: true
      }
    })
    wx.requestSubscribeMessage({
      tmplIds: [
        'aWyms3hhza3uUAH2aIK9vkUFiBeE5i4SiONpPRpwP7k',
        '0qEe1qVslIJy-wkRoa80QQaL8uqz_CU2HPP9_QAK3q8'
      ],
      success: function(res) {
        if (res['aWyms3hhza3uUAH2aIK9vkUFiBeE5i4SiONpPRpwP7k'] == 'accept') {
          wx.request({
            method: 'POST',
            url: `https://dw.dongwm.com/api/game/${game.id}/subscribe`,
            data: {},
            header: {
              "x-corran-token": getApp().globalData.openid
            },
            success: function({
              statusCode
            }) {
              if (statusCode == 200) {
                homePage.setData({
                  subscribeButton: {
                    text: '取消预定',
                    className: 'subscribed-button',
                    loading: false,
                    disabled: false
                  },
                  games: homePage.data.games.map(g => {
                    if (g.id == game.id) {
                      return {
                        ...game,
                        subscribed: true
                      }
                    } else {
                      return g
                    }
                  })
                })
              }
            }
          })
        } else {
          homePage.setData({
            subscribeButton: {
              ...homePage.data.subscribeButton,
              text: '确认预定',
              loading: false,
              disabled: false
            }
          })
        }
      }
    })

  }
})