import "./Checkbox.css"

interface Checkbox {
  checked: string;
  label: string;
  onChange: React.ChangeEventHandler
}

const Checkbox = ({checked, onChange, label}: Checkbox) => {
  return (
    <div className="checkbox-component">
      <div className="checkbox-wrapper-4">
        <input
          className="inp-cbx"
          id="morning"
          type="checkbox"
          value={checked}
          onChange={onChange}
        />
        <label className="cbx" htmlFor="morning">
          <span>
            <svg className="svg-1">
              <use xlinkHref="#check-4"></use>
            </svg>
          </span>
          <span>{label}</span>
        </label>
        <svg className="inline-svg">
          <symbol id="check-4" viewBox="0 0 12 10">
            <polyline points="1.5 6 4.5 9 10.5 1"></polyline>
          </symbol>
        </svg>
      </div>
    </div>
  );
};

export default Checkbox;