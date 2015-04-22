{{define "modifydish"}}
<form enctype="multipart/form-data" action="/admin/controller/dish" method="POST"> 
  <div class="row">
    <div class="col-md-4">
        <div class="form-group">
            <input type="text" class="form-control" name="name" placeholder="菜品名称" value="{{.Dish.Name}}">
        </div>
        <div class="form-group">
            <input type="text" class="form-control" name="price" placeholder="参考价格" value="{{.Dish.Price}}">
        </div>
        <div class="form-group">
            <label for="picurl">上传图片</label>
            <img src="">{{.Dish.Picurl}}
            <input type="file" name="picurl">
        </div>
    </div>
    <div class="col-md-4">
        <div class="form-group">
            <div class="input-group">
                <input type="text" class="form-control" id="datetimepicker" name="updatedtime" value="{{.Dish.Updatedtime}}" aria-describedby="basic-addon1" readonly>
                <span class="input-group-addon" id="basic-addon1"><i class="icon-th"></i></span>
            </div>
        </div>
        <div class="form-group">
            <div class="input-group">
              <select name="status">
                    <option value="true">选择状态</option>
                    <option value="true">显示</option>
                    <option value="false">隐藏</option>
              </select>
            </div>
        </div>
    </div>
  </div>

  <div class="row">
      <div class="col-md-8">
        <label for="content">简 介</label>
        {{template "editor" .}} 
        <textarea id="texteditor" name="synopsis" style="visibility:hidden"></textarea>
      </div>
      <input type="text" name="id" value="{{.Dish.Id}}" style="visibility:hidden">
  </div>          

  <div class="next">
    <button type="submit" id="moddish" class="btn btn-primary">提 交</button>
    <a type="button" class="btn btn-default" href="/admin/manage">取 消</a>
  </div>
</form>
<script type="text/javascript">
$('#datetimepicker').datetimepicker({
    fromat:'yyyy-mm-dd',
    minView:'month',
    language:'zh-CN',
    todayBtn:'linked',
    autoclose:'true',
});
$('#adddish').click(function(event) {
  var content = $('#editor').html();
  $('#texteditor').val(content);
});
</script>
{{end}}