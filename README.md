**Joining tables in postgresql**

* Queries can access multiple tables at once, or access the same table in such a way that multiple rows of the table are being processed at the same time. 
* A query that accesses multiple rows of the same or different tables at one time is called a join query
```
SELECT *
    FROM weather, cities
    WHERE city = name;
```

    
