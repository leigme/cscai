{{define "dish"}}
<div class="table-responsive">
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>名称</th>
        <th>创建时间</th>
        <th>修改时间</th>
        <th>参考价格</th>
        <th>点击</th>
        <th>操作</th>
      </tr>
    </thead>
    <tbody>
    {{range .Dishes}}
      <tr>
        <td>{{.Id}}</td>
        <td>{{.Name}}</td>
        <td>{{.Createdtime}}</td>
        <td>{{.Updatedtime}}</td>
        <td>{{.Price}}元</td>
        <td>{{.Click}}</td>
        <td><a type="button" class="btn btn-primary btn-xs" href="/admin/controller?op=moddish&id={{.Id}}">修改</a> / <a type="button" class="btn btn-danger btn-xs" href="/admin/controller?op=deldish&id={{.Id}}">删除</a></td>
      </tr>
      {{end}}
    </tbody>
  </table>
</div>
{{end}}