package singleton

type Manager interface {
    AddArticle(article *article.Article) error
    // Add other methods
}

type manager struct {
    db *gorm.DB
}

var Mgr Manager

func init() {
    db, err := gorm.Open("sqlite3", "../articles.db")
    if err != nil {
        log.Fatal("Failed to init db:", err)
    }
    Mgr = &manager{db: db}
}

func (mgr *manager) AddArticle(article *article.Article) (err error) {
    mgr.db.Create(article)
    if errs := mgr.db.GetErrors(); len(errs) > 0 {
        err = errs[0]
    }
    return
}
