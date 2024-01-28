from pathlib import Path

from setuptools import find_packages, setup


def read_file(path: Path | str) -> str:
    with Path(path).open('r') as f:
        return f.read().strip()


setup(
    name='pygo-tools',
    version='0.1.1',
    packages=find_packages(),
    include_package_data=True,
    entry_points={'console_scripts': ['build-ffi=setup_util.build_ffi:main']},
    python_requires='~=3.12',
    install_requires=['setuptools', 'cffi', 'wheel'],
    author='Rajan Khullar',
    author_email='rkhullar03@gmail.com',
    url='https://github.com/rkhullar/python-libraries/tree/main/pygo-tools',
    license='MIT NON-AI',
    long_description=read_file('readme.md'),
    long_description_content_type='text/markdown',
)
